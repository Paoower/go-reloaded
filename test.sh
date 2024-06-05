#!/bin/bash

INPUT_FILE="sample.txt"
OUTPUT_FILE="result.txt"
GO_PROGRAM="go run ."

verify_output() {
  expected_output="$1"
  printf "$expected_output" > expected_output.txt
  printf "\033[36m$expected_output\033[0m\n"
  if diff -q expected_output.txt $OUTPUT_FILE >/dev/null; then
    printf "\033[32mTest succeeded!\033[0m\n"
  else
    printf "\033[31mTest failed!\033[0m\n"
  fi
  rm expected_output.txt
}

print_io() {
    printf "\033[35m$(cat $INPUT_FILE)\033[0m\n"
    printf "\033[33m$(cat $OUTPUT_FILE)\033[0m\n"
}

printf "\n////////////////////////TESTS///////////////////////////"
printf "\n\033[35m◼\033[0m : Input \033[33m◼\033[0m : Output \033[36m◼\033[0m : Expected output"
printf "\nTest 1 :\n"
printf "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.\n" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."

printf "//////////////////////////////////////////////////////////\n"
printf "Test 2 :\n"
printf "Simply add 42 (hex) and 10 (bin) and you will see the result is 68." > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Simply add 66 and 2 and you will see the result is 68."

printf "//////////////////////////////////////////////////////////\n"
printf "Test 3 :\n"
printf "There is no greater agony than bearing a untold story inside you." > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "There is no greater agony than bearing an untold story inside you."

printf "//////////////////////////////////////////////////////////\n"
printf "Test 4 :\n"
printf "Punctuation tests are ... kinda boring ,don't you think !?" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Punctuation tests are... kinda boring, don't you think!?"

printf "//////////////////////////////////////////////////////////\n"
printf "Test 5 :\n"
printf "This sentence specifically test (cap, 3) for arguments in flags (up, 3) . DO YOU THINK YOU (low, 4) GET IT RIGHT ?" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "This Sentence Specifically Test for ARGUMENTS IN FLAGS. do you think you GET IT RIGHT?"

printf "//////////////////////////////////////////////////////////\n"
printf "Test 6 :\n"
printf "Here is a ' difficult one ', because you need to 'manage perfectly ' the quotes, and yes ' some times' it is ' better erase all your code and start again'." > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Here is a 'difficult one' , because you need to 'manage perfectly' the quotes, and yes 'some times' it is 'better erase all your code and start again' ."

printf "//////////////////////////////////////////////////////////\n"
printf "Test 7 :\n"
printf "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"

printf "//////////////////////////////////////////////////////////\n"
printf "Test 8 :\n"
printf "I have' to pack' 101 (bin) outfits. Packed 1a (hex) just to be sure" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "I have 'to pack' 5 outfits. Packed 26 just to be sure"

printf "//////////////////////////////////////////////////////////\n"
printf "Test 9 :\n"
printf "Don not be sad ,because     sad backwards is das . And das not good" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Don not be sad, because sad backwards is das. And das not good"

printf "//////////////////////////////////////////////////////////\n"
printf "Test 10 :\n"
printf "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '" > $INPUT_FILE
$GO_PROGRAM $INPUT_FILE $OUTPUT_FILE
print_io
verify_output "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"


