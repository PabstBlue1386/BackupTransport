package main

import (
	"fmt"
	"os/exec"
	"time"
)

const (
	ftpAdress   = "192.168.202.51:33344"
	ftpUser     = "ftpchiftec"
	ftpPassword = "BsZwPgCBmE9asvX"
	wcpPath     = "C:\\Program Files (x86)\\WinSCP\\WinSCP.com"
)

func main() {

	today := time.Now().Format("2006_01_02")
	localUtPath := "F:\\FTP\\BackupSQL\\UT-5"
	localRcPath := "F:\\FTP\\BackupSQL\\RC23"

	currentRcFileName := fmt.Sprintf("RC23_backup_%s_*.bak", today)
	currentUtFileName := fmt.Sprintf("UT_5_backup_%s_*.bak", today)

	remoteUtPath := "\\UT-day\\"
	remoteRcPath := "\\RC-day\\"

	command := ftm.Sprintf("
		open ftp://%s:%s@%s/ -passive=0
		 
	
	
	
	")

}

func copyFiles(remoteDir, localDir, fileName string) {

	fmt.Printf("Копирование %s из %s в %s...\n", localDir, remoteDir, fileName)

	ftpURL := fmt.Sprintf("ftp://%s:%s@%s/ -passive=0", ftpUser, ftpPassword, ftpAdress)
	cmd := exec.Command(wcpPath, ftpURL)

}
