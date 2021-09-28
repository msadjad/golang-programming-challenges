/*
 This is a solution for the LeetCode September Challenge. The link to the problem can be found bellow:
 https://leetcode.com/explore/challenge/card/september-leetcoding-challenge-2021/639/week-4-september-22nd-september-28th/3989/
*/
package main

import (
    "fmt"
    "strings"
)

func main() {
    numberOfUnique := numUniqueEmails([]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"})
    fmt.Println(numberOfUnique)
}

func numUniqueEmails(emails []string) int {
    uniqueEmails := map[string]bool{}
    
    for _, email := range emails {
        emailSplited := strings.Split(email, "@")
        beforePlus := strings.Split(emailSplited[0], "+")
        removedDots := strings.Join(strings.Split(beforePlus[0], "."), "")
        uniqueEmails[removedDots + "@" + emailSplited[1]] = true
    }

    return len(uniqueEmails)
}