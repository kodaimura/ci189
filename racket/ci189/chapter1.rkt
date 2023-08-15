#lang racket

;1.1
(define char-unique?
  (lambda (s)
    (let loop [[ls (string->list s)]]
      (cond
        ((null? ls) #f)
        ((member (car ls) (cdr ls)) #t)
        (else (loop (cdr ls)))))))

(not (char-unique? "abcd"))
(char-unique? "abad")
(char-unique? "abcb")

;1.2
(define permutations?
  (lambda (s1 s2)
    (string=? (list->string (sort (string->list s1) char<?))
              (list->string (sort (string->list s2) char<?)))))

(not (permutations? "abcdef" "abcdeg"))
(permutations? "abcdef" "afedcb")
(permutations? "abcdab" "aabdcb")

;1.3
(define replace
  (lambda (s from to)
    (let [[x (string-length from)]]
      (let loop [[ret ""] [s s]]
        (cond
          ((string=? s from) (string-append ret to))
          ((>= x (string-length s)) (string-append ret s))
          ((string=? (substring s 0 x) from) (loop (string-append ret to) (substring s x)))
          (else (loop (string-append ret (substring s 0 1)) (substring s 1))))))))

(define replace-space-with%20
  (lambda (s n)
    (replace (substring s 0 n) " " "%20")))

(replace"a b c" " " "%20")
(replace"abcabc" "ab" "%20")
(replace-space-with%20 "a b c" 4)
(replace-space-with%20 "a b c   " 6)

;1.4
(define palindrome?
  (lambda (s)
    (let [[s (replace s " " "")]]
      (string=? s (list->string (reverse (string->list s)))))))

(define palindrome-permutations?
  (lambda (s1 s2)
    (and (palindrome? s1)
         (palindrome? s2)
         (permutations? s1 s2))))

(palindrome? "abc dba")
(palindrome? "abc cba")
(palindrome? "ab c cba")
(palindrome-permutations? "taco cat" "atco ota")
(palindrome-permutations? "taco cat" "atco cta")

;1.5
;(string-length s1) = (string-length s2) の場合
(define can-convert-one-shot-replace?
  (lambda (s1 s2)
    (let loop [[count 0] [s1 s1] [s2 s2]]
      (cond
        ((> count 1) #f)
        ((string=? s1 "") #t)
        ((string=? (substring s1 0 1) (substring s2 0 1))
         (loop count (substring s1 1) (substring s2 1)))
        (else (loop (+ count 1) (substring s1 1) (substring s2 1)))))))

;(string-length s1) > (string-length s2) の場合
(define can-convert-one-shot-delete?
  (lambda (s1 s2)
    (cond
      ((string=? s2 "") (= 1 (string-length s1)))
      ((string=? (substring s1 0 1) (substring s2 0 1))
       (can-convert-one-shot-delete? (substring s1 1) (substring s2 1)))
      (else (string=? (substring s1 1) s2)))))

;(string-length s1) < (string-length s2) の場合
(define can-convert-one-shot-insert?
  (lambda (s1 s2)
    (can-convert-one-shot-delete? s2 s1)))

(define can-convert-one-shot?
  (lambda (s1 s2)
    (let [[s1l (string-length s1)]
          [s2l (string-length s2)]]
      (cond
        ((string=? s1 s2) #t)
        ((= s1l s2l) (can-convert-one-shot-replace? s1 s2))
        ((> s1l s2l) (can-convert-one-shot-delete? s1 s2))
        (else (can-convert-one-shot-insert? s1 s2))))))

(can-convert-one-shot? "pale" "pale")
(can-convert-one-shot? "pale" "ple")
(can-convert-one-shot? "pales" "pale")
(can-convert-one-shot? "pale" "bale")
(can-convert-one-shot? "pale" "bake")