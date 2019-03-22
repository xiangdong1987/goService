package model

import (
	"strings"
)

type AcNode struct {
	data     int32
	isEnd    bool
	Children []*AcNode
	length   int
	fail     *AcNode
}

type AcTrie struct {
	Root            *AcNode
	Dictionary      map[int32]int
	DicLength       int
	DictionaryLevel map[string]string
}

//初始化字典
func (ac *AcTrie) InitDictionary(wordList map[string]string) {
	i := 0
	ac.DictionaryLevel = wordList
	for value, _ := range wordList {
		for _, c := range value {
			if _, ok := ac.Dictionary[c]; ok {
				continue
			} else {
				ac.Dictionary[c] = i
				i++
				continue
			}
		}
	}
	ac.DicLength = i
}

//取出切片第一个
func pop(list []*AcNode) (*AcNode, []*AcNode) {
	if len(list) > 0 {
		a := list[0]
		b := list[1:]
		return a, b
	} else {
		return &AcNode{}, list
	}
}

//推入切片
func push(list []*AcNode, value *AcNode) []*AcNode {
	result := append(list, value)
	return result
}

//构建trie树
func (ac *AcTrie) AddWord(word string) {
	//fmt.Println(ac.Dictionary)
	nowNode := ac.Root
	i := 1
	for _, c := range word {
		if nowNode.Children[ac.Dictionary[c]] != nil {
			nowNode = nowNode.Children[ac.Dictionary[c]]
		} else {
			newNode := &AcNode{}
			newNode.Children = make([]*AcNode, ac.DicLength)
			newNode.data = c
			nowNode.Children[ac.Dictionary[c]] = newNode
			nowNode = newNode
		}
		if i == len([]rune(word)) {
			//fmt.Println(i)
			nowNode.isEnd = true
			nowNode.length = i
		}
		i++
	}
}

//初始化错误指针
func (ac *AcTrie) InitFailPoint() {
	var queue []*AcNode
	queue = push(queue, ac.Root)
	var p *AcNode
	for len(queue) > 0 {
		p, queue = pop(queue)
		for i := 0; i < len(p.Children); i++ {
			pc := p.Children[i]
			if pc == nil {
				continue
			}
			if pc == ac.Root {
				p.fail = ac.Root
			} else {
				q := p.fail
				for q != nil {
					qc := q.Children[ac.Dictionary[pc.data]]
					if qc != nil {
						pc.fail = qc
						break
					}
					q = q.fail
				}
				if q == nil {
					pc.fail = ac.Root
				}
			}
			push(queue, pc)
		}
	}
}

func (ac *AcTrie) Match(str string) (string, string) {
	p := ac.Root
	i := 1
	nowLevel := "0"
	var result []string
	for _, c := range str {
		//先判断字符是否在词典中
		if _, ok := ac.Dictionary[c]; !ok {
			//从头再来
			i++
			p = ac.Root
			continue
		}
		for p.Children[ac.Dictionary[c]] == nil && p != ac.Root {
			p = p.fail
			if p == nil {
				break
			}
		}
		if p == nil {
			//从头再来
			p = ac.Root
		}
		p = p.Children[ac.Dictionary[c]]
		tmp := p
		for tmp != ac.Root && tmp != nil {
			if tmp.isEnd == true {
				//fmt.Println(string(tmp.data))
				pos := i - tmp.length
				word := string([]rune(str)[pos : pos+tmp.length])
				//fmt.Println(word)
				tLevel := ac.DictionaryLevel[word]
				if tLevel > nowLevel {
					nowLevel = tLevel
				}
				result = append(result, word)
				str = strings.Replace(str, word, "***", -1)
				p = ac.Root
				//fmt.Println("Word is mach, pos is", pos, "length is", tmp.length)
			}
			tmp = tmp.fail
		}
		i++
	}
	return nowLevel, str
}
