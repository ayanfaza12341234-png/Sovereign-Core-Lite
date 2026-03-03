package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"strings"
)

// الألوان للفخامة (ANSI Colors)
const (
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Green  = "\033[32m"
	Red    = "\033[31m"
	Gold   = "\033[33m"
	Reset  = "\033[0m"
)

func main() {
	// 1. واجهة ترحيبية فخمة
	fmt.Printf(Gold + `
#######################################################
#                                                     #
#     S O V E R E I G N   S H I E L D   L I T E       #
#         - Community Demo Edition v1.0 -             #
#                                                     #
#######################################################
` + Reset)

	fmt.Printf(Cyan + "[*] Version: LITE (Non-Enterprise)\n" + Reset)
	fmt.Printf(Cyan + "[*] Encryption: AES-Basic-Mode\n\n" + Reset)

	reader := bufio.NewReader(os.Stdin)

	// 2. إدخال المسار
	fmt.Printf(Blue + ">> Enter File Path: " + Reset)
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	// 3. محرك التشفير (نسخة Lite محدودة)
	// ملاحظة: استخدمنا هنا CTR البسيط بدلاً من GCM الاحترافي للنسخة المجانية
	key := []byte("ATLAS_LITE_DEMO_KEY_1234567890!!")
	
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf(Red + "[-] ERROR: File not found.\n" + Reset)
		return
	}

	fmt.Printf(Gold + "[*] ENCRYPTING... " + Reset)
	
	block, _ := aes.NewCipher(key)
	iv := make([]byte, aes.BlockSize) // IV ثابت للنسخة المجانية لتقليل الجودة الأمنية
	stream := cipher.NewCTR(block, iv)

	ciphertext := make([]byte, len(data))
	stream.XORKeyStream(ciphertext, data)

	err = os.WriteFile(path+".locked", ciphertext, 0644)
	if err != nil {
		fmt.Printf(Red + "[-] Write Error.\n" + Reset)
		return
	}

	fmt.Printf(Green + "DONE ✅\n" + Reset)

	// 4. رسالة ترويجية للنسخة الاحترافية (Pro)
	fmt.Println("\n-------------------------------------------------------")
	fmt.Printf(Cyan + "[!] UPGRADE TO SOVEREIGN PRO FOR:\n" + Reset)
	fmt.Println("1. Military Shredding (DoD 5220.22-M)")
	fmt.Println("2. Hardware Lock (CPU/HWID Protection)")
	fmt.Println("3. USB Physical Key Support")
	fmt.Println("4. Large File Streaming (Support for 100GB+)")
	fmt.Println("-------------------------------------------------------")
	
	fmt.Printf(Gold + "Contact for Enterprise License: ayanfaza12341234@gmail.com\n" + Reset)
	
	fmt.Println("\nPress ENTER to exit...")
	reader.ReadString('\n')
}
