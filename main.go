package main

import (
	"fmt"
	"strings"
)

func ascii(word string) {
	// Define the ASCII art representation of each letter
	asciiArt := map[rune]string{
		'A': `
	  AA  
	 A  A 
	 AAAAA 
	 A  A 
	 A  A `,
		'B': `
	 BBB  
	 B  B 
	 BBB  
	 B  B 
	 BBB  `,
		'C': `
	  CCC 
	 C    
	 C    
	 C    
	  CCC `,
		'D': `
	 DDD  
	 D  D 
	 D  D 
	 D  D 
	 DDD  `,
		'E': `
	 EEEE 
	 E    
	 EEE  
	 E    
	 EEEE `,
		'F': `
	 FFFF 
	 F    
	 FFF  
	 F    
	 F    `,
		'G': `
	  GG  
	 G    
	 GGG  
	 G  G 
	  GG  `,
		'H': `
	 H  H 
	 H  H 
	 HHHH 
	 H  H 
	 H  H `,
		'I': `
	 III  
	  I   
	  I   
	  I   
	 III  `,
		'J': `
		J 
		J 
		J 
	 J  J 
	  JJ  `,
		'K': `
	 K  K 
	 K K  
	 KK   
	 K K  
	 K  K `,
		'L': `
	 L    
	 L    
	 L    
	 L    
	 LLLL `,
		'M': `
	 M     M 
	 MM M MM
	 M  M  M 
	 M     M 
	 M     M `,
		'N': `
	 N  N 
	 NN N 
	 N NN 
	 N  N 
	 N  N `,
		'O': `
	  OO  
	 O  O 
	 O  O 
	 O  O 
	  OO  `,
		'P': `
	 PPP  
	 P  P 
	 PPP  
	 P    
	 P    `,
		'Q': `
	  QQ  
	 Q  Q 
	 Q  Q 
	 Q QQ 
	  QQ Q`,
		'R': `
	 RRR  
	 R  R 
	 RRR  
	 R R  
	 R  R `,
		'S': `
	  SS  
	 S    
	  SS  
		S 
	  SS  `,
		'T': `
	 TTTT 
	  T   
	  T   
	  T   
	  T   `,
		'U': `
	 U  U 
	 U  U 
	 U  U 
	 U  U 
	  UU  `,
		'V': `
	 V   V
	 V   V
	 V   V
	  V V 
	   V  `,
		'W': `
	 W   W
	 W   W
	 W W W
	 WW WW
	 W   W`,
		'X': `
	 X   X
	  X X 
	   X  
	  X X 
	 X   X`,
		'Y': `
	 Y   Y
	  Y Y 
	   Y  
	   Y  
	   Y  `,
		'Z': `
	 ZZZZ 
		Z 
	   Z  
	  Z   
	 ZZZZ `,
		' ': `
		  
		  
		  
		  
		  `,
	}

	// Split each ASCII art into lines
	lines := make([][]string, len(word))
	for i, char := range strings.ToUpper(word) {
		if ascii, ok := asciiArt[char]; ok {
			asciiLines := strings.Split(ascii, "\n")
			lines[i] = asciiLines
		} else {
			fmt.Println("Letter not found:", string(char))
		}
	}

	// Print ASCII art horizontally
	for i := 0; i < len(lines[0]); i++ {
		for j := 0; j < len(word); j++ {
			if lines[j] != nil {
				fmt.Print(lines[j][i])
			}
		}
		fmt.Println()
	}

}

func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scan(&word)

	ascii(word)
}
