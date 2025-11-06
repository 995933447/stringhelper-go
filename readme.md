### 字符串助手函数集合

### 支持函数:
Snake(s string) string 字符串下转划线风格,如输入AbCde,获得ab_cde

Camel(s string) string 字符串转驼峰风格,如输入Ab_cd_ef,获得AbCdEf

LowerFirstASCII(s string) string 字符串首字母小写

UpperFirstASCII(s string) string 字符串首字母大写

FindAllSubstringIdx(s, sub string) []int 获取字符串s中所有子串sub的下标

GenRandomStr(length int64, mod uint32) string 生成随机字符串,length是最终随机字符串的长度,
mod代表字符串风格,支持以下常量定义三种风格:
````
const (
    RandomStringModNumberPlusLetter           = 1 // 1-9数字+52个大小写字母进行随机拼接
    RandomStringModNumberPlusLetterPlusSymbol = 2 // 1-9数字+52个大小写字母+~!@#$%^&*(){}|\[]?/进行随机拼接
    RandomStringModNumber                     = 3 // 1-9随机拼接
)
````

