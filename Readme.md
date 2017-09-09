### 哥伦布编码  解析sps

* 参考: https://github.com/nareix/bits

```
unsigned integer = 9
binary = 1001 (4 bits)
exp-golomb code = 000 1 010 (7 bits)
```
```
解析过程如下: 9 = 2的3次幂 - 1 + bytevalue(010)
= 7 + 2
3次幂的3是第一次遇见"1"bit时的index值,然后读取index个bit
```