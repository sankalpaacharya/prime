package lexer

import "prime/token"

type Lexer struct {
  input string 
  position int
  nextPosition int
  char byte
}

func (l *Lexer) readChar(){
  if l.nextPosition>=len(l.input){
    l.char = 0 
  }
  else{
    l.char = l.input[l.position]
  }
  l.position = l.nextPosition
  l.nextPosition+=1
}

func Next(input string){
  l:= &Lexer{input:input}
  l.readChar()
  return l
}






