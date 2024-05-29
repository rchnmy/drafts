package partyrobot // https://exercism.org/tracks/go/exercises/party-robot/solutions/rchnmy

import (
    "strings"
    "strconv"
)

var str strings.Builder

// Welcome greets a person by name.
func Welcome(name string) string {
    defer str.Reset()
  
    msg := []string{"Welcome to my party, ", name, "!"}
    for _, m := range msg {
        str.WriteString(m)
    }
    return str.String()
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    defer str.Reset()
  
    sage := strconv.FormatInt(int64(age), 10)
    msg := []string{"Happy birthday ", name, "! ", "You are now ", sage, " years old!"}
    for _, m := range msg {
        str.WriteString(m)
    }
    return str.String()
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
    defer str.Reset()
    
    null := "0"
    stable := strconv.FormatInt(int64(table), 10)
    for len(stable) < 3 {
          stable = null + stable
      }

    rdist := strconv.FormatFloat(distance, 'f', 1, 64)
    msg := []string{"Welcome to my party, ", name, "!\nYou have been assigned to table ", stable, ". Your table is ", direction, ", exactly ", rdist, " meters from here.\nYou will be sitting next to ", neighbor, "."}
    for _, m := range msg {
        str.WriteString(m)
    }
    return str.String()
}
