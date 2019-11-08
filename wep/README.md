# wep stands for the word error possibility

in this time, to calculate wep,


1. calculate the maximizing likelihood decoding rule
2. calculate wep for each given error values in float


# Given condition

input code | codeword
-----------|----------
00|000
01|101
10|110
11|111



# Result

1. The maximum likelihood decoding rule

Lambda - codeword(encoded) | included received word in lambda
---------------------------|-----------------------------------
lambda1 - C1(000) | 000
lambda2 - C2(101) | 001, 101
lambda3 - C3(110) | 010, 100, 110
lambda4 - C4(111) | 011, 111

2. WEP for e = {0.0, 0.1, 0.2, 0.3, 0.4, 0.5}

![image](https://user-images.githubusercontent.com/29854277/68441870-d38b2580-0212-11ea-87b5-faa29ef2d741.png)
