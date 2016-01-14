# Longest Compund Word Finder

## Installation
Get the code  
```go get github.com/mantishK/quiz```  
Run it  
```cd $GOPATH/src/github.com/mantishK/quiz```  
```go run main.go -file="/path/to/words.list"```

## Solution
#### Storing the words in space and time effecient way  
Build a tree for each starting character of the words. Add each character to the node of the respective tree and mark the node containing the last character of each node as "end".
#### Retreiving the words
Traverse the tree to match each character and if the node marked "end" is reached at the last character of the word, then we have a match.
#### Finding compound words
For each word, begin with the first character and test if it exists and if the rest of the sub-word is a compound word recursively.