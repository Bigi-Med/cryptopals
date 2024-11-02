package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main(){
    //var valHex string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    //var val64 string = convertHexToBase64(valHex)
    //fmt.Printf(val64)

    //var val1 string = "1c0111001f010100061a024b53535009181c"
    //var val2 string = "686974207468652062756c6c277320657965"
    //var xored string = xor(val1,val2)
    //fmt.Printf(xored)
    //text,key,_ := xorWithAllBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    //fmt.Printf(text)
    //fmt.Printf(string(key))
    path := "decode.txt"
    lines := readFile(path)
    text := detectSingleCharXor(lines)
    fmt.Printf(text)

}

func convertHexToBase64(valHex string) string {
    decodedHex,_ := hex.DecodeString(valHex)
    val64 := base64.StdEncoding.EncodeToString(decodedHex)
    return val64
}

func xor(val1, val2 string) string {
    val1H,_ := hex.DecodeString(val1)
    val2H,_ := hex.DecodeString(val2)
    b1 := []byte(val1H)
    b2 := []byte(val2H)
    result := make([]byte,len(b1))
    for i := range b1 {
     result[i] = b1[i]^b2[i]   
    }
    var xored string = hex.EncodeToString(result)
    return xored
}

func scoreText(msg []byte) int{
    score :=0
    for _,b := range msg {
        if unicode.IsLetter(rune(b)) || b == ' '{
         score++   
        }
    }
    return score

}

func xorWithAllBytes(codeHex string) (string,byte,int) {
 //convert Hex input to string
 code,_ := hex.DecodeString(codeHex)
 bcode := []byte(code)
 var bestScore int
 var bestKey byte
 var bestText string

 for key := byte(0); key<255;key++{
     decrypted := make([]byte, len(code))
     for i := range len(bcode) {
      decrypted[i] = bcode[i]^key  
     }

     score := scoreText(decrypted)

     if score > bestScore {
         bestScore = score  
         bestText = string(decrypted)
         bestKey = key
     }
 }

 return bestText,bestKey,bestScore
 
}

func readFile(path string) []string{
    content,_ := os.ReadFile(path)
    lines := strings.Split(string(content),"\n")
    return lines
}

func detectSingleCharXor(line []string) string {
    var bestScore int
    var bestText string
   for i := 0 ; i<len(line);i++{
       text,_,score := xorWithAllBytes(line[i])
       if(score > bestScore){
        bestScore = score
        bestText = text
       }
   }
   return bestText
}














