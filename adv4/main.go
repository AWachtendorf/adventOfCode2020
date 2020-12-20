package adv4

import (
	"fmt"
	"strconv"
	"strings"
)

//puzzleInput from website
const puzzleInput = "byr:1985\neyr:2021 iyr:2011 hgt:175cm pid:163069444 hcl:#18171d\n\neyr:2023\nhcl:#cfa07d ecl:blu hgt:169cm pid:494407412 byr:1936\n\necl:zzz\neyr:2036 hgt:109 hcl:#623a2f iyr:1997 byr:2029\ncid:169 pid:170290956\n\nhcl:#18171d ecl:oth\npid:266824158 hgt:168cm byr:1992 eyr:2021\n\nbyr:1932 ecl:hzl pid:284313291 iyr:2017 hcl:#efcc98\neyr:2024 hgt:184cm\n\niyr:2017 pid:359621042\ncid:239 eyr:2025 ecl:blu byr:1986 hgt:188cm\n\neyr:2027 hgt:185cm hcl:#373b34 pid:807766874 iyr:2015 byr:1955\necl:hzl\n\niyr:2017 hcl:#7d3b0c hgt:174cm\nbyr:1942 eyr:2025 ecl:blu pid:424955675\n\neyr:2026 byr:1950 hcl:#ceb3a1\nhgt:182cm\niyr:2016 pid:440353084 ecl:amb\n\nhcl:a4c546\niyr:1932 pid:156cm eyr:2034 hgt:193 ecl:zzz byr:2025\n\nhcl:#ceb3a1 eyr:2020 pid:348696077 hgt:163cm\necl:hzl\nbyr:1921 iyr:2016\n\necl:gmt eyr:2031 iyr:2018 byr:1971 hgt:152in pid:454492414\nhcl:z\n\nhcl:#341e13 byr:1921 iyr:2020\npid:072379782 eyr:2022 hgt:166cm cid:253 ecl:brn\n\necl:blu hgt:75in byr:1954 eyr:2026 iyr:2012 hcl:#623a2f pid:328598886\n\nbyr:2004 eyr:2035 hcl:#7d3b0c pid:359128744 iyr:2020 hgt:65cm\necl:#70f23f\n\neyr:1988\npid:171cm byr:2003\niyr:1984\ncid:50\nhcl:z hgt:66cm ecl:#7a4c6e\n\npid:9632440323 eyr:1964 hgt:63cm\necl:#fab0c5 hcl:z iyr:1945 byr:1986\n\npid:936403762 ecl:#337357 byr:1997\ncid:196 iyr:2020\neyr:2030 hgt:165cm\nhcl:#7d3b0c\n\nbyr:1931 pid:488791624 hgt:169cm ecl:blu\neyr:2029 hcl:#fffffd iyr:2013\n\nhcl:#733820 hgt:76in pid:517689823\neyr:2028 byr:1988\necl:brn iyr:2016\n\neyr:2023 hcl:#fffffd hgt:190cm iyr:2015 ecl:brn pid:739536900 byr:1951\n\necl:brn\nbyr:1986 cid:262 hcl:#efcc98 pid:880203213 hgt:185cm iyr:2018 eyr:2029\n\npid:181cm hgt:113 hcl:z ecl:#2c2d2c iyr:1961 byr:2021 eyr:2031\n\nhcl:#ceb3a1 iyr:2020\nbyr:1977\nhgt:192cm\npid:338237458 eyr:2030 ecl:amb\n\niyr:1953 byr:2025 hgt:66cm eyr:1932\npid:181cm\necl:#6f0b15 hcl:f79cb7\ncid:109\n\nhcl:#6b5442 pid:164cm ecl:blu\nhgt:176cm byr:2015\niyr:2010 eyr:2029\n\neyr:2035\npid:085002665 ecl:#f88074 iyr:2018 hcl:#602927\nhgt:169cm\n\nbyr:1958\nhcl:z\npid:0468194841 iyr:2016 eyr:2007\nhgt:152cm\necl:#1c7a89 cid:124\n\nhcl:z pid:233430735 byr:2021 eyr:2026\niyr:1953 ecl:#64769d hgt:184\n\nhgt:70cm pid:156397147\niyr:2014 ecl:#d6ada0\nbyr:2030\nhcl:#cfa07d\n\necl:amb\nbyr:1990\niyr:2017 hgt:164cm hcl:10f33a\ncid:293 eyr:2020 pid:332276985\n\npid:163252726 eyr:2026\nhgt:163cm\niyr:2011 hcl:#efcc98\necl:hzl byr:1936\n\nhgt:157cm iyr:2019 pid:078770050 hcl:#efcc98 byr:1967 eyr:2030\necl:gry cid:190\n\nhgt:184cm ecl:amb pid:851379559 hcl:#ceb3a1 byr:1946 eyr:2022\niyr:2017 cid:280\n\nhgt:171cm byr:1942 pid:830156471 hcl:#cfa07d ecl:gry eyr:2032\niyr:2022\n\nbyr:2013 ecl:#67cbe8 eyr:2024\npid:242908367\nhgt:76cm\niyr:2025\nhcl:796bda\n\necl:amb iyr:2019\nbyr:1945 eyr:2021 hcl:#602927 pid:550065206\n\nhgt:72in ecl:brn byr:1956 pid:253685193 iyr:2017 eyr:2023\nhcl:#6b5442\n\neyr:2032 iyr:2019\nhgt:176cm\necl:oth pid:800237895 hcl:#888785 byr:1979\n\neyr:2026 iyr:2020 cid:226 pid:882830512\nhcl:#866857 byr:1929 ecl:amb\nhgt:60in\n\nhcl:#cfa07d ecl:oth\niyr:2015 pid:807837948 byr:1966 eyr:2030 hgt:191in\n\nbyr:1969 iyr:2012 eyr:2024\ncid:244 ecl:hzl hcl:#18171d pid:344160556\n\neyr:2020 pid:718422803\nhcl:#18171d\nhgt:181cm\nbyr:1925 ecl:amb\niyr:2019\n\nbyr:1943 pid:740807220 hgt:72in ecl:amb\niyr:2013 eyr:2022\nhcl:#cfa07d\n\nhcl:#733820\nbyr:1986 iyr:2016 hgt:184cm cid:333\npid:768188726 ecl:oth eyr:2030\n\neyr:2022 byr:1996 hcl:#341e13 ecl:hzl iyr:2015 hgt:160cm\npid:516401532\n\nhgt:182cm ecl:grn pid:336742028 iyr:2014 hcl:#34f021 byr:1967\neyr:2029\n\nbyr:2030\nhgt:142 iyr:2029 eyr:2040 hcl:426fc5\ncid:312\npid:169cm\necl:#069ff7\n\nhgt:169cm ecl:gry hcl:#6b5442 iyr:2012 byr:1949 pid:131835020 eyr:2022\n\nhgt:70cm iyr:2012\neyr:2037\nhcl:64fd76\ncid:175 pid:4880649770 ecl:grn byr:2029\n\niyr:2013 hcl:#7d3b0c eyr:2024 hgt:190cm pid:659772377 cid:226 ecl:oth byr:1958\n\necl:lzr hgt:163cm pid:013605217\nbyr:2000\neyr:2020\nhcl:z iyr:2024\n\ncid:131 pid:896076106\nhcl:#c0946f byr:1930\nhgt:162cm eyr:2023 ecl:oth iyr:2017\n\nbyr:1935 iyr:2012\npid:942509879\necl:amb\nhgt:185cm cid:152 eyr:2024 hcl:#866857\n\necl:#e490a3 hcl:4813a2 hgt:176cm pid:778369210 iyr:2020\neyr:2035 byr:2020\n\nbyr:2006 ecl:amb pid:148409219\nhgt:189cm\neyr:2021 hcl:z iyr:2028\n\nhgt:188in hcl:#9ed525\niyr:2018 ecl:grn eyr:2021\npid:065515632 byr:2012\n\ncid:109 hgt:167cm\npid:545112664 ecl:grn hcl:#a62fea eyr:2026\niyr:2012\nbyr:1921\n\npid:174997024\niyr:2012\neyr:2030\necl:grn\nhgt:150cm\nbyr:1997\nhcl:#866857\n\npid:451921339\nhgt:181cm\nhcl:#888785 iyr:2017 eyr:2026 byr:1936\necl:hzl\n\nhgt:187in\nhcl:#866857 ecl:grn pid:623919686 eyr:2028 iyr:2011\nbyr:2016\n\nbyr:2001\necl:gry eyr:2023 pid:324948416\nhcl:ef16f8 cid:139 hgt:184in iyr:2026\n\nbyr:1954 hcl:#341e13 eyr:2023 pid:129321944 iyr:2012\nhgt:183cm\necl:amb\n\nhgt:164cm pid:596870080\necl:hzl eyr:2021 iyr:2017 hcl:#a97842\nbyr:1951\n\niyr:2013 byr:1944 hcl:#cfa07d\nhgt:168cm cid:72 pid:160531632\necl:grn\n\niyr:2012 pid:900043442 hcl:#ceb3a1 cid:124 byr:1941\necl:blu hgt:156cm\neyr:2025\n\neyr:2021 hgt:61in iyr:2020 ecl:grn byr:1933\n\nbyr:1971 cid:175\neyr:2028 hcl:#efcc98 iyr:2013 hgt:170cm\npid:225213589\n\npid:147112660 hcl:#ceb3a1 eyr:2029 hgt:159cm ecl:grn iyr:2014\nbyr:1967\n\niyr:2015 pid:502975636 hgt:71in byr:1994\nhcl:#18171d ecl:amb eyr:2029\n\nbyr:1948 hcl:#b6652a hgt:171in pid:181cm iyr:2019 ecl:grt cid:87\n\npid:859849571 ecl:amb hcl:#6b5442\nhgt:193cm byr:1980\niyr:2017\neyr:2020\n\ncid:125 pid:508147848\nhcl:06ea75 iyr:1997 byr:2010 ecl:#c707f7 eyr:1970 hgt:161\n\neyr:2020 cid:326 byr:1989 ecl:gry hgt:160cm hcl:#cc080c pid:319135853 iyr:2010\n\necl:utc\npid:531595917 hgt:180cm byr:1987\neyr:2024 hcl:#cfa07d iyr:2025\n\necl:gry byr:2007\neyr:2028\niyr:2025\npid:6072964414 hgt:59cm hcl:#888785\n\npid:791025828 ecl:hzl hgt:178cm\niyr:2017\nhcl:#733820\nbyr:1960 eyr:2021 cid:66\n\nbyr:1991 iyr:1934\ncid:304 hgt:183cm ecl:grn\npid:408294229\neyr:2027 hcl:#623a2f\n\necl:blu hgt:181cm eyr:2024 iyr:2010\npid:633234602 hcl:#2ce009\nbyr:1985\n\nhcl:#c0946f hgt:192cm\niyr:2012 pid:120684397 ecl:grn eyr:2027\nbyr:1974\n\neyr:2026\npid:068304960 hgt:190cm byr:1925 iyr:2020 ecl:oth\n\nhcl:#733820\nhgt:168cm cid:307 iyr:2014 byr:1981 ecl:hzl pid:898831724 eyr:2026\n\nhgt:73cm\neyr:2038\nbyr:1980 ecl:gry iyr:2027 pid:678846912 hcl:z\n\nhgt:150cm cid:261 eyr:2021\nhcl:z pid:159cm iyr:2014 ecl:hzl\nbyr:1955\n\npid:#172650 ecl:gry eyr:2040 hcl:z iyr:2013 hgt:169cm byr:2008 cid:290\n\niyr:2017 byr:1998\nhcl:#ceb3a1 pid:274178898 eyr:2027 ecl:brn\nhgt:183cm\n\neyr:2024 cid:183 ecl:grn\nbyr:1946\nhgt:63in hcl:#6b5442 iyr:2017\n\nhgt:97 byr:1990\niyr:2019\necl:grn\npid:587580330\nhcl:#341e13 eyr:2022\n\necl:oth\npid:441517075 hcl:#c0946f iyr:2015 hgt:188cm eyr:2024 byr:1920\n\nhgt:191in pid:185cm iyr:1993\nhcl:93033d\neyr:2034 ecl:dne\n\npid:591478424 ecl:grn hcl:#888785\nbyr:1929 eyr:2023 hgt:173cm iyr:2017\n\niyr:1954\nhgt:63cm\nhcl:bdf2e0 ecl:amb pid:#912f46\n\nbyr:1956 iyr:2012 hgt:73in pid:986643426\necl:blu\ncid:235 eyr:2025\n\nhcl:#cfa07d\ncid:320 byr:1930\nhgt:172cm\necl:oth eyr:2024 iyr:2019\n\nbyr:1935 hgt:182cm pid:22794407 hcl:1b96fb eyr:1961 iyr:1941 ecl:#5e80cd cid:70\n\niyr:2020 eyr:2021\necl:amb\nhgt:59in pid:594829025 hcl:#93092e\nbyr:1976\n\nhcl:#a97842 eyr:2030\nbyr:1937 iyr:2018 cid:295 ecl:oth\nhgt:166cm pid:282634012\n\nhgt:171cm hcl:#623a2f byr:1956\npid:068178613 cid:214\niyr:2012 eyr:2026 ecl:brn\n\nbyr:1921\nhgt:161cm hcl:#888785\necl:brn pid:010348794\neyr:2023 iyr:2011\n\nhcl:#a97842 iyr:2010\nbyr:1955 eyr:2024\npid:473791166\necl:brn\nhgt:175cm\n\neyr:2028 ecl:grn pid:186196675 byr:1945 hgt:155cm cid:349\niyr:2011 hcl:#6b5442\n\nhgt:161cm eyr:2030 cid:221\npid:994494879 hcl:#733820 iyr:2012 ecl:blu\nbyr:1957\n\neyr:1993 iyr:2022 hcl:z byr:2020 pid:013428192 hgt:62cm\necl:dne\n\nhgt:178cm eyr:2029 hcl:#733820 byr:1962 iyr:2017 ecl:blu pid:567713232\n\nhcl:#fffffd\nbyr:1928 pid:390162554\neyr:2030 cid:79 hgt:150cm ecl:amb iyr:2019\n\neyr:2030 cid:320 hgt:171cm hcl:#888785 pid:540720799 ecl:amb iyr:2012 byr:1979\n\nbyr:1921\necl:oth pid:204986110 eyr:2023 hgt:154cm iyr:2017 hcl:#341e13 cid:126\n\neyr:2020 cid:175 ecl:dne byr:1983 iyr:2016 hcl:#c0946f hgt:65cm\n\nhgt:191cm\niyr:2010 cid:295 byr:1984 eyr:2025 hcl:#cfa07d pid:799775698\necl:amb\n\niyr:2020 cid:278 hcl:#c0946f byr:1970 pid:773144393 eyr:2024 hgt:180cm\n\nhgt:176cm\nbyr:1963\npid:252396293 iyr:2012 ecl:brn hcl:#ceb3a1\neyr:2030\n\npid:545130492\nbyr:2030 iyr:2020\nhgt:190cm eyr:2034 ecl:blu hcl:#fffffd\n\nhcl:#a97842 pid:032201787 hgt:190cm ecl:gry\neyr:2028 iyr:2012 byr:1994\n\nhcl:#a97842 pid:064591809\necl:hzl byr:1927 hgt:165cm\niyr:2011\neyr:2028\n\ncid:77\nbyr:2005\nhgt:125 iyr:1923 ecl:#605d73\neyr:2022 pid:90184674 hcl:z\n\ncid:301 pid:106820988\niyr:2018\nhcl:#cfa07d eyr:2029\nbyr:1993\nhgt:193cm ecl:grn\n\nhcl:#623a2f\ncid:118\necl:oth pid:75827285\nhgt:189cm iyr:2010\neyr:2030 byr:1976\n\necl:blu iyr:2023 eyr:1996\nhgt:66cm cid:251 byr:1972 hcl:z\npid:557774244\n\nbyr:2002\nhgt:169cm pid:629420566 eyr:2026 ecl:grn hcl:#341e13\ncid:166 iyr:2019\n\niyr:2026 hcl:9b83a1 eyr:1979\necl:dne hgt:111 pid:176cm\n\npid:#89718c byr:2026\nhcl:2ca5c7 hgt:142 eyr:2040\necl:lzr iyr:2029\n\necl:grn\nbyr:2022 eyr:2020\npid:7024869 hgt:123 iyr:2019 hcl:z\n\nhcl:#733820 hgt:155cm ecl:grn iyr:2020 byr:1955 eyr:2028\npid:217362007\n\nhcl:#18171d ecl:gry\nbyr:1971 hgt:193cm\neyr:2020\npid:352009857 iyr:2013\n\nbyr:2018\nhgt:175in ecl:xry iyr:2015\neyr:2036\ncid:171 pid:6132398 hcl:#efcc98\n\npid:839955293\nbyr:1928 hcl:#fffffd ecl:hzl iyr:2011\nhgt:162cm eyr:2023\n\nhgt:175cm pid:482827478 eyr:2028\nhcl:#6b5442 ecl:blu byr:1932 iyr:2010\n\niyr:2020 hcl:#866857\necl:brn byr:1933 cid:269 pid:003931873 hgt:188cm\neyr:2022\n\nbyr:1981 hcl:#fffffd hgt:160cm cid:311 ecl:brn eyr:2025\npid:930857758 iyr:2014\n\nhcl:#cfa07d hgt:73in\necl:gry\npid:383281251\niyr:2013 byr:1934 eyr:2026\n\nbyr:1988 eyr:2026 pid:458002476\niyr:2017\nhgt:175cm ecl:amb\n\neyr:1987\nbyr:2020 pid:299341304\nhcl:#341e13 iyr:1935 cid:125\nhgt:168cm\necl:gry\n\niyr:2014 hcl:#b6652a pid:445799347\nhgt:188cm byr:1960\neyr:2030 cid:290 ecl:amb\n\neyr:2023\nhgt:75cm hcl:#733820 cid:195 byr:1933\necl:amb pid:062770586 iyr:2019\n\nhgt:168cm\neyr:2021\npid:725299968 ecl:grn byr:1932\niyr:2016 hcl:#888785\n\nhgt:161cm hcl:#ceb3a1 byr:1962 eyr:2026 iyr:2013 ecl:amb pid:695426469 cid:227\n\necl:dne hcl:#ceb3a1 iyr:2013 eyr:2022\npid:434786988 byr:1956 hgt:183cm\n\npid:697500517\nbyr:1968 hgt:169cm hcl:#fffffd ecl:grn cid:143\niyr:2010\neyr:2027\n\nbyr:2029 ecl:amb hgt:175in iyr:2015 hcl:#ceb3a1\npid:39839448\neyr:2021 cid:105\n\npid:0985607981 ecl:hzl iyr:2012\neyr:2021 byr:2024 hcl:5cad22\nhgt:190cm\n\nhcl:#b6652a hgt:178cm cid:222 byr:1992 ecl:grn\niyr:2011 pid:419544742\n\niyr:2019 byr:1960 ecl:hzl eyr:2021 hgt:184cm cid:66 hcl:#866857 pid:412920622\n\neyr:2025 hcl:#888785 iyr:2018 byr:1956 pid:698098389 ecl:grn hgt:173cm\n\necl:blu byr:1935\npid:354892542 hgt:161cm\niyr:2018\neyr:2021 hcl:#b6652a\n\necl:oth cid:287 iyr:2028 byr:1953 eyr:2027 hcl:#7d3b0c hgt:151cm\npid:211411839\n\niyr:2018 byr:1934 hcl:#a97842\npid:859748861\necl:oth hgt:175cm eyr:2025\n\nbyr:1930 iyr:2018 eyr:2022\nhgt:175cm\nhcl:#292092\necl:brn pid:987163365\n\nhgt:167in hcl:#888785 eyr:2040 pid:4646402867 byr:2013 iyr:1941 ecl:#389aec\n\necl:hzl hcl:#602927\nhgt:168cm eyr:2026\ncid:235 iyr:2016\nbyr:1942\n\niyr:1975 pid:11337832 ecl:#a25273 hgt:151 byr:2017\n\neyr:1979\nhgt:71cm\nbyr:2003 hcl:7e7da7 pid:151cm ecl:#a8afb3 iyr:1937\n\neyr:2021 hgt:74in hcl:#cfa07d iyr:2014 byr:1932\npid:641867677 ecl:grn\n\necl:gry hgt:185cm pid:556229206 iyr:2013\nbyr:1984\nhcl:#fffffd eyr:2028\n\neyr:2020 byr:1989\necl:grn pid:618876158 hcl:z\nhgt:176cm iyr:2025\n\neyr:2025 byr:2001 hcl:#cdb7f9\npid:377402126 ecl:hzl hgt:184cm iyr:2019\n\nbyr:1939 hgt:180cm eyr:2029 ecl:oth hcl:#733820 iyr:2016\npid:733456875\n\npid:883743276\nhcl:#7d3b0c eyr:2022 ecl:blu\nbyr:1928 hgt:150cm cid:150 iyr:2013\n\nhgt:60cm ecl:#43f03d eyr:1994 byr:1975\niyr:1980 pid:169cm\n\nhgt:104 byr:2029 eyr:2040 hcl:64a9b2\npid:83898860\niyr:1990\necl:#938bbe\n\npid:284399238 ecl:gry hcl:#888785 iyr:2019 hgt:168cm byr:1944\neyr:2022\n\nhcl:#733820 pid:486515752 ecl:grn hgt:188in byr:1941 iyr:2017 eyr:2005\n\niyr:2010\nbyr:1978 hgt:160cm eyr:2003\necl:oth\nhcl:#efcc98 pid:584668011\n\nbyr:1944 ecl:gry pid:962700562 iyr:2011 hcl:#866857 eyr:2022\nhgt:191cm\n\nhcl:z pid:758583213 iyr:1941 ecl:gry eyr:2007\nhgt:67 byr:2022\ncid:215\n\nbyr:1988\necl:#ae2a9b hcl:#fe9d14 iyr:2012\npid:411550516 hgt:169cm eyr:2038\n\npid:400034647 byr:1927 hgt:165cm\niyr:2017 ecl:brn eyr:2024 cid:144 hcl:#341e13\n\nhcl:#733820 hgt:153cm eyr:2027\nbyr:1935 pid:217121064 cid:120 iyr:2012 ecl:grn\n\nhgt:168cm hcl:#866857 iyr:2012 pid:1527348755\nbyr:1946 eyr:2028 cid:184 ecl:amb\n\nhcl:#a97842\nbyr:1967\nhgt:152cm eyr:2030\necl:blu\npid:929661915 iyr:2018\n\npid:671485026\nhgt:188cm byr:1974 iyr:2015 ecl:grn cid:268 eyr:2021 hcl:#c0946f\n\npid:789877199 iyr:2011 cid:219 eyr:2029\necl:oth byr:1991\nhcl:#866857 hgt:154cm\n\ncid:137 pid:059579902\neyr:2020 byr:1952\nhcl:#18171d iyr:2020\nhgt:172cm ecl:oth\n\npid:182cm iyr:1997 byr:2012\neyr:2034\nhgt:161in ecl:#528abf hcl:b7d2fe\n\nhgt:192cm ecl:oth iyr:2017 pid:264538307 byr:1994 cid:285\nhcl:#18171d eyr:2030\n\nhcl:#efcc98\npid:38036608\neyr:2010\niyr:2026\nbyr:2027\ncid:239 ecl:zzz hgt:74\n\niyr:2012\neyr:2022 hgt:178cm\nhcl:#888785\necl:hzl\nbyr:1998 pid:000080585\n\npid:719620152 hcl:#b6652a cid:133\necl:hzl\nbyr:1983 iyr:2012 hgt:175cm\neyr:2024\n\ncid:155 eyr:1977 iyr:2019 ecl:#28de8b byr:1941 hcl:#602927 hgt:173cm pid:493773064\n\niyr:2010\npid:842124616 ecl:hzl eyr:2025 cid:146 hcl:#733820 hgt:166cm byr:1987\n\nhcl:fd4dcf byr:2006 iyr:2011 pid:820797708 eyr:2020 hgt:189cm\necl:gry\n\niyr:1971 pid:22107293 hcl:#5b3f01 cid:257\necl:hzl\nhgt:60cm eyr:2000 byr:1965\n\nbyr:1932 eyr:2028\nhcl:#6b5442 ecl:amb pid:947149686\niyr:2015 hgt:187cm\n\nhcl:#a97842\ncid:260\nhgt:167cm eyr:2027 byr:1973 ecl:oth pid:741678753 iyr:2016\n\npid:334234443 ecl:gry hcl:#18171d eyr:2020\niyr:2016 hgt:159cm byr:1926\n\nhgt:118 eyr:1929 iyr:2013\npid:987139064\ncid:196\nhcl:#cfa07d ecl:#f72601 byr:1929\n\nbyr:1924\npid:623185744 iyr:2012 cid:341 hcl:#602927 hgt:192cm eyr:2022\n\niyr:2012 byr:1971 hgt:168cm cid:146 pid:673038025 hcl:#866857 eyr:2020 ecl:hzl\n\neyr:2023 iyr:2017\npid:205596613 cid:298 hcl:#341e13\nhgt:169cm ecl:oth\nbyr:1996\n\necl:blu pid:775831730\neyr:2029 iyr:1924 hgt:168cm hcl:z\n\nbyr:2023 hgt:181cm\npid:4365105095 iyr:2021\necl:lzr eyr:2024 hcl:z\n\nhgt:184cm byr:1987 pid:175cm ecl:#83a5fa eyr:2023\n\neyr:2021 pid:422371422 ecl:oth iyr:2015 hcl:#866857\nbyr:1963 hgt:174cm\n\npid:006970943\nhcl:#2f22ef iyr:2020\necl:gry\nbyr:1922\neyr:2024 hgt:163cm\n\ncid:160 byr:2015\neyr:2038 hcl:z ecl:grt hgt:166 iyr:2026\npid:#14978f\n\nhgt:178cm eyr:2021 iyr:2016 pid:471529794\nhcl:#b6652a cid:192\necl:grn byr:1970\n\niyr:2015 ecl:brn hcl:#602927 hgt:187cm\npid:729284172\neyr:2024 byr:1932\n\ncid:153\necl:dne eyr:2005\npid:178cm iyr:2028\nbyr:2029 hgt:160in hcl:482a92\n\nbyr:1995 iyr:2012 hcl:#866857 hgt:159cm\neyr:1950 ecl:gry pid:183cm\n\npid:875885919\nhgt:159cm\niyr:2011\necl:gry byr:1988 hcl:#341e13 eyr:2028\n\npid:2390267705 hcl:#7d3b0c byr:2009\neyr:2017 ecl:grn hgt:183cm iyr:2015\n\necl:brn eyr:2029 hcl:#866857 iyr:2020 hgt:180cm byr:2001\npid:668021168\n\nhcl:#c0946f\neyr:2024 ecl:amb pid:013487714 byr:1965 hgt:172cm cid:320 iyr:2020\n\neyr:2025 pid:115479767 hcl:#866857 ecl:oth\nhgt:163cm iyr:2010 byr:1999\n\nbyr:1967 iyr:2011 cid:112 hcl:#733820\neyr:2040 ecl:grt\nhgt:66 pid:804536366\n\nhgt:163 pid:1764836278 eyr:2035\niyr:2021\nhcl:z ecl:#f1bb27\n\nhcl:#efcc98 hgt:176cm byr:1994 pid:590539278 ecl:grn iyr:2011 eyr:2021\n\niyr:2017 eyr:2024 hgt:167cm hcl:#b62e29 pid:495674801\nbyr:1970 ecl:brn\n\nhgt:168cm pid:993244641\nbyr:1968\neyr:1926\nhcl:#b6652a ecl:brn\niyr:2023\n\nhgt:63in hcl:z pid:594070517\neyr:2021 ecl:oth\niyr:2017\nbyr:2000\n\neyr:2030 pid:272955042 cid:319 iyr:2011 ecl:amb byr:1999 hcl:#888785 hgt:158cm\n\neyr:2025\npid:814305816 byr:1945 ecl:brn hgt:162cm iyr:2018\nhcl:#a97842\ncid:229\n\nbyr:1996 eyr:2026 pid:582584802 hcl:#c0946f iyr:2020 ecl:grn\nhgt:162cm\n\neyr:2027\nhgt:155cm byr:1925\nhcl:#888785\ncid:182\niyr:2014 ecl:brn\npid:250884352\n\nhgt:173cm cid:135\niyr:2017 pid:661330507 byr:1950 eyr:2020 ecl:gry hcl:#18171d\n\npid:208932950\neyr:2030 hgt:179cm\niyr:2013\necl:oth\nbyr:1981\ncid:58 hcl:#6b5442\n\nhcl:#f183e7 iyr:2014\nhgt:159cm pid:614579850 ecl:gry eyr:2029\ncid:186 byr:1962\n\neyr:2027 hcl:#db3405 byr:1938 pid:194516631 cid:167 hgt:177cm ecl:oth\n\nhgt:68in hcl:#733820 pid:228644594 eyr:2030 ecl:gry iyr:2010 cid:334 byr:1951\n\niyr:2017 hcl:#341e13\npid:#6a28c9 hgt:154cm ecl:gry\nbyr:1966 eyr:2023\n\npid:250155574 cid:84\nhgt:157cm ecl:grn byr:1937 iyr:2017 eyr:2024 hcl:#b6652a\n\npid:831823039 eyr:2028 iyr:2015 ecl:gry\nhgt:192cm cid:137 byr:1922\nhcl:#6b5442\n\nhgt:193cm byr:1941 eyr:2024 cid:56\nhcl:#623a2f ecl:amb\npid:351293754 iyr:2016\n\nbyr:1947 iyr:2012 ecl:hzl hcl:#602927 eyr:2028 pid:252010138 hgt:152cm\n\nhcl:#a97842 pid:801192586 ecl:hzl iyr:2018 hgt:193cm byr:1928 cid:323\neyr:2028\n\nhgt:151cm\npid:756347561 ecl:hzl\neyr:2024 cid:161\niyr:2016 hcl:#623a2f\nbyr:2002\n\npid:648012871 iyr:2015 ecl:blu\neyr:2025 hcl:#623a2f byr:1973 hgt:177cm\n\nbyr:1999 hcl:#ceb3a1 cid:345 eyr:2025 ecl:#b29a96 pid:093304949\niyr:2017 hgt:93\n\nhcl:#b6652a\niyr:2018 ecl:grn\nbyr:1951 pid:077278028 eyr:2024 hgt:62in\n\nhgt:164cm pid:410770618 byr:1958\niyr:2019\neyr:2030\necl:gry hcl:#fffffd cid:293\n\necl:grt\neyr:2039\nhcl:z pid:188cm byr:2022\niyr:2027 hgt:76cm\n\necl:grn iyr:2012 hgt:150cm eyr:2024\nbyr:1926 pid:954310029 cid:64\nhcl:#fffffd\n\necl:oth eyr:2027 pid:091152959 hgt:180cm hcl:#ceb3a1 iyr:2015 cid:350\nbyr:1924\n\niyr:2017 hcl:#49a793 eyr:2021 cid:144 byr:1966\npid:717543257\nhgt:161cm\necl:hzl\n\neyr:2025 ecl:brn hgt:60in pid:391973520 byr:1928 cid:77\niyr:2012\nhcl:#602927\n\niyr:2013 hgt:161cm pid:784483994 byr:1991\nhcl:#cfa07d\neyr:2024 ecl:grn\n\necl:hzl iyr:1967 byr:2009 cid:265 hgt:180in pid:168cm\neyr:1966\n\neyr:2024 iyr:2019 pid:534453983\nbyr:2028 ecl:oth hcl:#341e13 hgt:193cm\n\neyr:2029 iyr:2010 hcl:#623a2f ecl:gry hgt:152cm pid:572128647\nbyr:1996\n\niyr:2014 byr:1981 cid:176\necl:grn hgt:183cm pid:974469723 eyr:2027\n\neyr:2029 pid:233353682 byr:1968\necl:gry hgt:181cm iyr:2011\nhcl:#efcc98\n\nhgt:61 iyr:2005 cid:203 ecl:gmt pid:157cm hcl:z\nbyr:2013\n\niyr:2020\nbyr:1923 ecl:blu eyr:2026 pid:069770502 hgt:69cm\nhcl:z\n\nbyr:1997 hgt:160cm\nhcl:z iyr:2021 eyr:1920 pid:9374226872\n\necl:hzl eyr:2024 pid:537492791 hgt:186cm byr:1952\nhcl:#cfa07d\niyr:2020\n\nhgt:73cm byr:1974\necl:xry iyr:2016 cid:133\nhcl:e741f5 pid:186cm\n\npid:161cm\nbyr:1950\neyr:2028 ecl:hzl hcl:#7d3b0c\niyr:2014 hgt:158cm\n\necl:#2c491e\nhcl:f8fe13 byr:2022\nhgt:137 iyr:1948\neyr:2040 pid:#959a0f\n\nbyr:1923 hgt:70in\npid:904825661 hcl:#b6652a iyr:2010 eyr:2020\necl:oth\n\niyr:2013\necl:blu pid:858020233 byr:1950 hgt:61in\n\nhcl:#18171d\niyr:2016\necl:amb pid:613754206 byr:1975 hgt:164cm eyr:2025\n\nbyr:1938\niyr:2017 hcl:#623a2f cid:191 eyr:2027 hgt:174cm pid:287108745 ecl:amb\n\niyr:2025 hcl:#623a2f byr:2019 hgt:170cm\ncid:233 pid:55323151 ecl:amb eyr:2037\n\necl:amb\nhgt:177cm hcl:#b6a3ce eyr:2025 byr:1967 pid:506927066\niyr:2018 cid:93\n\nbyr:1964 hgt:173cm eyr:2030 cid:106 pid:587635596 iyr:2012\nhcl:#fb5993\necl:hzl\n\necl:lzr pid:190cm hcl:44746d eyr:1955 hgt:66cm iyr:1990 byr:2003\n\necl:brn byr:1968 cid:216 hgt:181in hcl:#b6652a iyr:2016 eyr:2020 pid:0208311541\n\necl:hzl hgt:181cm\neyr:1977 byr:2018 pid:527754216 hcl:#c0946f\n\necl:grn hcl:#efcc98\nbyr:1935 eyr:2025 iyr:2018 hgt:65in pid:396444938 cid:293\n\nhgt:64in ecl:oth\nhcl:#18171d\npid:105602506 byr:1973\neyr:2022\niyr:2014\n\neyr:2039 hgt:64\necl:#ab45a8 byr:2009\niyr:2025 pid:182cm hcl:d1614a cid:103\n"

//valid passports for debugging
//const puzzleInput = "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980\nhcl:#623a2f\n\neyr:2029 ecl:blu cid:129 byr:1989\niyr:2014 pid:896056539 hcl:#a97842 hgt:165cm\n\nhcl:#888785\nhgt:164cm byr:2001 iyr:2015 cid:88\npid:545766238 ecl:hzl\neyr:2022\n\niyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"

//passPorts holds a string
type passPorts struct {
	passport string
}

//newPassPort creates and returns a passport
func newPassPort(s string) *passPorts {
	p := &passPorts{
		passport: s,
	}
	return p
}

//allPassPorts holds all passPorts in a slice
type allPassPorts struct {
	passports []*passPorts
}

//readAllData reads the data of the puzzle input and processes it
//to put it in a slice
func (a *allPassPorts) readAllData() *allPassPorts {
	//the puzzle input is processed.
	//first we split the string at the blank lines
	separatedString := strings.Split(puzzleInput, "\n\n")
	for _, j := range separatedString { //we replace all returns with blank spaces
		j = strings.ReplaceAll(j, "\n", " ")
		a.passports = append(a.passports, newPassPort(j)) //then we put the cleaned and separated strings in passports
	}
	return a
}

//singlePassport is also a passport, but holds a map
type passportMap struct {
	sPassports []map[string]string
}

//newMap returns a new map
func newMap() map[string]string {
	var passPortMap = map[string]string{}

	return passPortMap
}

//insertInMap takes a map and inserts key and value
func insertInMap(m map[string]string, key, val string) map[string]string {
	m[key] = val

	return m
}

//parseInMap parses the data in a passport.
//the slice data is one long string, we need to split it and process it into a map
func (a *allPassPorts) parseInMap() *passportMap {
	s := &passportMap{} //
	for _, j := range a.passports {
		split := strings.ReplaceAll(j.passport, ":", " ") //every entry is split
		split2 := strings.Fields(split) //then we read the values in a slice
		//println(fmt.Sprintf("%v", split2))
		tempMap := newMap() //we need new map
		for i := range split2 {
			if i == len(split2)-1 {
				break
			}
			if i%2 == 0 { //
				insertInMap(tempMap, split2[i], split2[i+1])
			}
		}
		s.sPassports = append(s.sPassports, tempMap)
	}
	return s
}

//checkValidation counts how many passports are valid
func (s *passportMap) checkValidation() int {
	checksum := 0

	for _, j := range s.sPassports {
		if checkInMap(j) {
			checksum++
		}
	}
	return checksum
}

//checkInMap checks first if all fields are present, afterwards if the values correspond to the policy
func checkInMap(m map[string]string) bool {
	byr, validbyr := m["byr"]
	iyr, validiyr := m["iyr"]
	eyr, valideyr := m["eyr"]
	hgt, validhgt := m["hgt"]
	hcl, validhcl := m["hcl"]
	ecl, validecl := m["ecl"]
	pid, validpid := m["pid"]

	if validbyr && validiyr && valideyr && validhgt && validhcl && validecl && validpid {
		println("\npassport contains all fields") // prints "Freddy exists"

		if checkByrIyrEyr(byr, 1920, 2002) &&
			checkByrIyrEyr(iyr, 2010, 2020) &&
			checkByrIyrEyr(eyr, 2020, 2030) &&
			checkHgt(hgt) &&
			checkHcl(hcl) &&
			checkEcl(ecl) &&
			checkPid(pid) {
			return true
		} else {

			return false
		}
	} else {
		println("\npassport doesn't contain all fields")
		return false
	}

}

func checkByrIyrEyr(tocheck string, min, max int) bool {
	if len(tocheck) < 4 || len(tocheck) > 4 {
		//println("\ndoesn't match charactercount")
		//println(fmt.Sprintf("\n%s is invalid", tocheck))
		return false
	}
	byrAsInt, err := strconv.Atoi(tocheck)
	if err != nil {
		println(err)
	}
	if byrAsInt >= min && byrAsInt <= max {
		//println("\nnot in range")
		//println(fmt.Sprintf("\n%s is invalid", tocheck))
		return true

	} else {
		//println(fmt.Sprintf("\n%s is valid", tocheck))
		return false
	}
}

func checkHgt(tocheck string) bool {
	if strings.Contains(tocheck, "cm") {
		tocheck = strings.ReplaceAll(tocheck, "cm", "")
		toInt, err := strconv.Atoi(tocheck)
		if err != nil {
			println(err)
		}
		if toInt >= 150 && toInt <= 193 {
			//println(fmt.Sprintf("\n%s is valid", tocheck))
			return true
		} else {
			//println(fmt.Sprintf("\n%s is invalid", tocheck))
			return false
		}

	}
	if strings.Contains(tocheck, "in") {
		tocheck = strings.ReplaceAll(tocheck, "in", "")
		toInt, err := strconv.Atoi(tocheck)
		if err != nil {
			println(err)
		}
		if toInt >= 59 && toInt <= 76 {
			//println(fmt.Sprintf("\n%s is valid", tocheck))
			return true
		} else {
			//println(fmt.Sprintf("\n%s is invalid", tocheck))
			return false
		}

	} else {
		//println(fmt.Sprintf("\n%s is invalid", tocheck))
		return false
	}
}

func checkHcl(tocheck string) bool {
	const invalid = "ghijklmnopqrstuvwxyz"

	if strings.Contains(tocheck, "#") && !strings.Contains(tocheck, invalid) && len(tocheck) == 7 {
		println(fmt.Sprintf("\n%s is valid", tocheck))
		return true
	} else {
		println(fmt.Sprintf("\n%s is invalid", tocheck))
		return false
	}
}

func checkEcl(tocheck string) bool {
	validcolor := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, j := range validcolor {
		if tocheck == j {
			//println(fmt.Sprintf("\ncolor %s is valid", tocheck))
			return true
		}
	}
	//println(fmt.Sprintf("\ncolor %s is invalid", tocheck))
	return false
}

func checkPid(tocheck string) bool {
	if len(tocheck) == 9 {
		//println(fmt.Sprintf("\nPID %s is valid with length %v", tocheck, len(tocheck)))
		return true
	} else {
		//println(fmt.Sprintf("\nPID %s is invalid with length %v", tocheck, len(tocheck)))
		return false
	}
}

func main() {
	a := allPassPorts{}
	s := &passportMap{}
	a.readAllData()
	s = a.parseInMap()
	checkSum := s.checkValidation()
	println(fmt.Sprintf("\n %v Passports are valid", checkSum))
}
