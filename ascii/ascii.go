package ascii

import (
	"fmt"
	"strings"
)

func AsciiRun() {
	fmt.Print("Enter text to convert to ASCII art: ")
	var text string
	_, _ = fmt.Scanln(&text)

	asciiArt := GenerateASCIIArt(text)
	fmt.Println(asciiArt)
}

func GenerateASCIIArt(text string) string {
	artMap := map[rune]string{
		'A': `  A  
 A A 
AAAAA
A   A
A   A`,
		'B': `BBBB  
B   B 
BBBB  
B   B 
BBBB  `,
		'C': ` CCCC 
C     
C     
C     
 CCCC `,
		'D': `DDDD  
D   D 
D   D 
D   D 
DDDD  `,
		'E': `EEEEE 
E     
EEEE  
E     
EEEEE `,
		'F': `FFFFF 
F     
FFFF  
F     
F     `,
		'G': ` GGG  
G     
G GGG 
G   G 
 GGG  `,
		'H': `H   H 
H   H 
HHHHH 
H   H 
H   H `,
		'I': `  I   
 II   
  I   
  I   
 III  `,
		'J': `   J  
   J  
   J  
   J  
JJJ   `,
		'K': `K   K 
K  K  
KKK   
K  K  
K   K `,
		'L': `L     
L     
L     
L     
LLLLL `,
		'M': `M   M 
MM MM 
M M M 
M   M 
M   M `,
		'N': `N   N 
NN  N 
N N N 
N  NN 
N   N `,
		'O': ` OOO  
O   O 
O   O 
O   O 
 OOO  `,
		'P': `PPPP  
P   P 
PPPP  
P     
P     `,
		'Q': ` QQQ  
Q   Q 
Q   Q 
Q  QQ 
 QQQ Q`,
		'R': `RRRR  
R   R 
RRRR  
R  R  
R   R `,
		'S': ` SSS  
S     
 SSS  
    S 
SSSS  `,
		'T': `TTTTT 
  T   
  T   
  T   
  T   `,
		'U': `U   U 
U   U 
U   U 
U   U 
 UUU  `,
		'V': `V   V 
V   V 
V   V 
 V V  
  V   `,
		'W': `W   W 
W   W 
W W W 
W W W 
W   W `,
		'X': `X   X 
X   X 
 X X  
  X   
X   X `,
		'Y': `Y   Y 
Y   Y 
 YYY  
  Y   
  Y   `,
		'Z': `ZZZZZ 
   Z  
  Z   
 Z    
ZZZZZ `,
		' ': `     
     
     
     
     `,
	}

	var asciiArt strings.Builder

	for i, char := range text {
		if i > 0 {
			// Add space between characters
			asciiArt.WriteString("\n\n")
		}

		// Convert the character to uppercase
		char = rune(strings.ToUpper(string(char))[0])
		if ascii, found := artMap[char]; found {
			asciiArt.WriteString(ascii)
		} else {
			asciiArt.WriteString("Character not found in ASCII art.")
		}
	}

	return asciiArt.String()
}
