A Dictionary of Kagome v2
===

A dictionary package of [kagome v2](http://github.com/ikawaha/kagome/tree/v2). This software includes a binary and/or source version of data from

* mecab-naist-jdic

which can be obtained from

* https://open-jtalk.sourceforge.net/

as a part of `openjtalk`'s source
  
 # Feature Fields
 
Features are information given to a word, such as follows:

```
こんな  名詞,形容動詞語幹,*,*,*,*,こんな,コンナ,コンナ,0/3,C2
寒い    形容詞,自立,*,*,形容詞・アウオ段,基本形,寒い,サムイ,サムイ,2/3,*
日      名詞,非自立,副詞可能,*,*,*,日,ヒ,ヒ,0/1,C3
で      助詞,格助詞,一般,*,*,*,で,デ,デ,1/1,動詞%F1
私      名詞,代名詞,一般,*,*,*,私,ワタシ,ワタシ,0/3,C3
は      助詞,係助詞,*,*,*,*,は,ハ,ワ,0/1,名詞%F1/動詞%F2@0/形容詞%F2@0
泳が    動詞,自立,*,*,五段・ガ行,未然形,泳ぐ,オヨガ,オヨガ,2/3,*
ない    助動詞,*,*,*,特殊・ナイ,基本形,ない,ナイ,ナイ,1/2,動詞%F3@0/形容詞%F2@1
よ      助詞,終助詞,*,*,*,*,よ,ヨ,ヨ,0/1,名詞%F1/動詞%F1/形容詞%F1
。      記号,句点,*,*,*,*,。,。,。,*/*,*
EOS
```
 
 |No.|feature|description|
 |:---|:---|:---|
 | 0| POS hierarchy | The POS name and each level in its hierarchical structure. |
 | 1| POS hierarchy 1 | |
 | 2| POS hierarchy 2 | |
 | 3| POS hierarchy 3 | |
 | 4| Inflectional Type| Inflection type indicates a category that is an inflected form, e.g. 五段・カ行促音便.|
 | 5| Inflectional Form| Inflected form, e.g. 連用タ接続. |
 | 6| Base Form| A form of dictionary headword, e.g. 行っ -> 行く. |
 | 7| Reading|A possible reading for an entry. Readings are given in katakana, e.g. コウエン.|
 | 8| Pronunciation| A possible pronunciations for an entry. Pronunciations are given in katakana, .e.g. コーエン. |
 | 9| Accent | A possible accent position for an entry, e.g. 3/5 |
 | 10| Accent Connection Type | Behavior of an entry when it connects with another entry, for more information, see [Unidic Manual](https://clrd.ninjal.ac.jp/unidic/UNIDIC_manual.pdf) |

# Licence

MIT

Please be noted that openjtalk is released under modified BSD licence, you may also want to check it in [NOTICE](./NOTICE.txt).