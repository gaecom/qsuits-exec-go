package qsuits

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"qsuits-exec-go/src/config"
	"qsuits-exec-go/src/utils"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func CheckJavaRuntime() (javaPath string, version string, err error) {
	javaPath, err = exec.LookPath("java")
	if err != nil {
		return javaPath, version, err
	}
	version, err = GetJavaVersion(javaPath)
	return javaPath, version, err
}

func GetJavaVersion(javaPath string) (version string, err error) {
	cmd := exec.Command(javaPath, "-version")
	var versionRet string
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return version, err
	} else {
		if out.Len() > 0 {
			versionRet = out.String()
		} else {
			versionRet = stderr.String()
		}
	}
	var versionFields []string
	versionFields = strings.Split(strings.Split(versionRet, "\n")[0], " ")
	version = strings.Trim(versionFields[2], "\"")
	return version, nil
}

func CheckJavaVersion(version string, minimum int) (err error) {
	items := strings.Split(version, ".")
	var ver int
	if strings.EqualFold(items[0], "1") {
		ver, err = strconv.Atoi(items[1])
	} else {
		ver, err = strconv.Atoi(items[0])
	}
	if err != nil {
		return err
	}
	if ver < minimum {
		err = errors.New(fmt.Sprintf("please update your java to jdk%d or above", minimum))
		return err
	}
	return nil
}

//func GetJavaPath(jdkPath string) (javaPath string) {
//	osName := runtime.GOOS
//	if strings.Contains(osName, "darwin") {
//		javaPath = filepath.Join(jdkPath, "bin", "java")
//		_, err := CheckJavaVersion(javaPath)
//		if err != nil {
//			javaPath = filepath.Join(jdkPath, "Contents", "Home", "bin", "java")
//		}
//	} else if strings.Contains(osName, "windows") {
//		javaPath = filepath.Join(jdkPath, "bin", "java.exe")
//	} else {
//		javaPath = jdkPath
//	}
//	return javaPath
//}

func SetJdkMod(path string, jdkPath string, minimum int) (javaPath string, err error) {

	if len(path) == 0 {
		return javaPath, errors.New("no valid path")
	}
	var version string
	var firstErr error
	osName := runtime.GOOS
	if strings.Contains(osName, "darwin") {
		javaPath = filepath.Join(jdkPath, "bin", "java")
		version, err = GetJavaVersion(javaPath)
		if err != nil {
			firstErr = err
			javaPath = filepath.Join(jdkPath, "Contents", "Home", "bin", "java")
			version, err = GetJavaVersion(javaPath)
		}
	} else if strings.Contains(osName, "windows") {
		javaPath = filepath.Join(jdkPath, "bin", "java.exe")
		version, err = GetJavaVersion(javaPath)
	} else {
		javaPath = filepath.Join(jdkPath, "bin", "java")
		version, err = GetJavaVersion(javaPath)
	}
	if err != nil {
		if firstErr != nil {
			err = errors.New(fmt.Sprintf("%s, %s", firstErr, err))
		}
		return javaPath, err
	}
	err = CheckJavaVersion(version, minimum)
	if err != nil {
		return javaPath, err
	}
	isExists, err := utils.FileExists(filepath.Join(path, ".qsuits"))
	if isExists && err != nil {
		// is directory
	} else if isExists && err == nil {
		err = errors.New(fmt.Sprintf("%s is already exists, but is only a file", filepath.Join(path, ".qsuits")))
		return javaPath, err
	} else {
		err = os.MkdirAll(filepath.Join(path, ".qsuits"), 0755)
		if err != nil {
			return javaPath, err
		}
	}
	modPath = filepath.Join(path, ".qsuits", "jdk.mod")
	modFile, err := os.Create(modPath)
	defer modFile.Close()
	if err != nil {
		return javaPath, err
	}
	size, err := modFile.WriteString(javaPath)
	if err != nil {
		return javaPath, err
	}
	if size <= 0 {
		return javaPath, errors.New("no content wrote")
	}
	return javaPath, nil
}

func GetJavaPathFromMod(path string) (javaPath string, err error) {

	if len(path) == 0 {
		return javaPath, errors.New("no valid path")
	}
	modPath = filepath.Join(path, ".qsuits", "jdk.mod")
	modFile, err := os.Open(modPath)
	defer modFile.Close()
	if err != nil {
		return javaPath, err
	}
	content, err := ioutil.ReadAll(modFile)
	if err != nil {
		return javaPath, err
	}
	javaPath = string(content)
	return javaPath, nil
}

func JdkDownload() (jdkFileName string, err error) {

	osName := runtime.GOOS
	osArch := runtime.GOARCH
	jdkFileName = "jdk-8u231"
	if strings.Contains(osName, "darwin") {
		jdkFileName += "-macosx-x64.dmg"
	} else if strings.Contains(osName, "linux") {
		if strings.Contains(osArch, "64") {
			jdkFileName += "-linux-x64.tar.gz"
		} else if strings.Contains(osArch, "86") {
			jdkFileName += "-linux-i586.tar.gz"
		} else {
			return jdkFileName, errors.New(fmt.Sprintf("no jdk tar.gz file to download for this os_arch: %s_%s",
				osName, osArch))
		}
	} else if strings.Contains(osName, "windows") {
		if strings.Contains(osArch, "64") {
			jdkFileName += "-windows-x64.exe"
		} else if strings.Contains(osArch, "86") {
			jdkFileName += "-windows-i586.exe"
		} else {
			return jdkFileName, errors.New(fmt.Sprintf("no jdk executable file to download for this os_arch: %s_%s",
				osName, osArch))
		}
	} else {
		return jdkFileName, errors.New(fmt.Sprintf("no jdk to download for this os_arch: %s_%s", osName, osArch))
	}

	jdkUrl := config.ADDRESS + jdkFileName
	done := make(chan struct{})
	go utils.SixDotLoopProgress(done, "jdk-downloading")
	err = ConcurrentDownloadWithRetry(jdkUrl, jdkFileName, 2097152, 30 * time.Second, 5)
	if err != nil {
		if strings.Contains(err.Error(), "certificate") || strings.Contains(err.Error(), "handshake") {
			jdkUrl = "http" + strings.TrimPrefix(jdkUrl, "https")
			err = ConcurrentDownloadWithRetry(jdkUrl, jdkFileName, 2097152, 30 * time.Second, 5)
		}
	}
	if err != nil {
		if strings.Contains(err.Error(), "request canceled") || strings.Contains(err.Error(), "copy error size") {
			err = ConcurrentDownloadWithRetry(jdkUrl, jdkFileName, 2097152, 0, 5)
		}
	}
	done <- struct{}{}
	close(done)
	if err != nil {
		return jdkFileName, err
	} else {
		fmt.Println(" -> succeed.")
		return jdkFileName, nil
	}
}
