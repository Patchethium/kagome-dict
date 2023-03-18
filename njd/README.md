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
公園	名詞,一般,*,*,*,*,公園,コウエン,コーエン
に	助詞,格助詞,一般,*,*,*,に,ニ,ニ
行っ	動詞,自立,*,*,五段・カ行促音便,連用タ接続,行く,イッ,イッ
た	助動詞,*,*,*,特殊・タ,基本形,た,タ,タ
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