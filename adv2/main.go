package main

import (
	"strconv"
	"strings"
)

//puzzleInput from website
const puzzleInput = "15-19 k: kkkkkkkkkkkkzkkkkkkk\n1-11 s: sbssswsqsssssrlss\n8-9 b: pbbbbbbkbz\n4-10 w: wwccwcqwdmbktjrxhw\n1-6 x: jvscgqsnt\n1-7 x: xxxxxxcx\n6-10 s: smssssfskssdwvtcss\n6-12 q: qqqqzqqjqfqdqq\n3-7 d: ddwbzbf\n12-14 s: ssdssssssssmsq\n4-11 w: wwwwprvgklvwtxwpwwww\n6-7 j: jjjjjjz\n1-6 l: xxlnll\n14-15 n: nnnnxnkmnnnnnvfnnnj\n5-6 h: hbhhhhz\n4-6 b: brbhbrb\n2-7 q: zmqszpqwcq\n9-15 z: zzzzzzzpczzzzzzz\n6-10 m: mmmmmmmmmmmm\n8-14 l: llllllbljlllzllz\n4-10 c: cbccjxhlcclpf\n1-4 d: mpvglsjhsddtsnjsdqw\n2-8 f: ffffffftfb\n4-14 q: qqcpqqqqqtzhqqqqknq\n3-5 m: nctdm\n1-4 l: lllg\n11-12 s: sssssspsssfvxks\n5-9 l: fvsllcvgsmlzw\n3-5 d: dddtpd\n3-11 x: xxjkxxcxqvlprxgx\n8-9 b: bbbbbbsbs\n6-8 f: zffffnffjs\n7-8 v: vvvxvvvvvv\n6-9 b: nbvrbptfbbnbxb\n8-9 c: cccccccrc\n14-15 v: vvlvvvvvvvvvvlvv\n1-5 f: lvndmpdf\n3-4 l: sllfl\n5-14 w: jknqvcwwzwgfwwww\n2-6 b: vbbcbc\n3-4 x: xxbwf\n12-16 h: hhhkhhhthxqnhzhhhn\n1-4 j: vjjjj\n14-17 m: mmmmmmmxmmbpmcmmmmz\n1-9 z: rzzzzzzzzzzzzzzh\n7-8 r: vrmrrnrrrrnr\n8-11 z: fzzzzzzzzzz\n3-5 l: kdjlljpllz\n11-15 h: xhhhhhbjhshhkhbhhhht\n4-11 h: qhbnhhhhhdtwhqsh\n8-10 r: xrrrrrrrrrr\n5-9 q: qqqqpgqqq\n6-14 g: kgvfhqvhggglrgddgg\n6-8 h: ssctfnhhvhxhhxphhlc\n4-18 h: hhhqhhhhqhkhhhhhcl\n2-6 t: pvttttttmtx\n1-12 b: gfhbwwbbsvbcfb\n8-10 m: fmmmmmmzmmrmkq\n16-20 h: hhlhhhhhhhhbhhsfhhhh\n4-8 q: xqtjqqqq\n5-12 g: xzrbngggggnngb\n4-5 j: jjjsj\n6-12 b: bxmbbbbbbbbrpbb\n13-14 n: nnnnnnnnnnnnng\n14-16 s: wfnqltmpsksrtvdc\n6-13 g: ggmggjggggggggg\n5-9 v: wvvvrvvdmv\n1-5 h: hhhhdvhh\n6-8 g: bggggcgg\n5-9 w: wjfwhhwwtwwt\n10-13 n: nnnnnnnnnnnnqnn\n1-3 f: qfff\n15-19 t: ttttttttttttmtstrttt\n1-2 d: bddd\n5-6 v: vvvvvn\n3-4 r: vhrprrgncrcrbkml\n12-14 c: ccccccfcccclccncg\n11-12 m: mpmmbmpxdtdp\n16-17 f: qsfffffjwfttfxffr\n7-13 f: fdjcfffffvffv\n7-9 z: zsnzzzpxz\n8-14 r: grrrrxrrrnjrrs\n6-8 f: svndffdh\n13-17 b: bbbbbbbbbbbbbbbbp\n3-14 f: fkxsfkvvxbfbhfff\n4-15 d: bdndvlwwnzsqdhd\n12-14 f: hfmffkmjlfffqf\n5-8 v: qmlctvvvvjtvj\n5-8 w: wwpxvrgwbfwwtbwlx\n11-13 f: fffqffffffffn\n2-5 s: wvmhsmswn\n12-14 h: hhhhhhhhhhhhhxh\n8-14 l: lllllllclcllllllql\n13-15 l: lllllvgsllsllllllfnl\n2-8 c: cjlccvcc\n1-5 l: plqllpsllnl\n1-5 k: kvfqkkk\n9-13 s: dpvfnsnmksfss\n3-4 b: mfgb\n1-8 j: wjjwjjjxzxn\n9-11 v: vvvvwvvvvvs\n7-8 v: vvrvvhhrhvpvvbv\n9-12 d: dscvhvlndzgf\n10-17 l: llqllpllllllllllgllx\n3-4 k: nkkgkv\n14-19 q: qrqqqqqqqqqqwzqfqqq\n9-11 c: nqccccccccncccccc\n5-19 n: pfndnknwrnnbnjqndnnt\n2-6 s: nsftszgk\n3-9 j: jjfjjjjjfj\n1-7 m: xmmtzmmmm\n5-13 j: jljrjhjlgjtrnq\n2-13 f: fmfzggfmffsth\n12-13 m: mmmmmmmmmmmtvm\n11-12 k: kkkkkkkkkkjk\n8-10 s: sssssssksks\n3-4 w: wfwwbwswd\n2-9 l: sllqllldll\n4-6 f: zfcfrf\n9-12 m: sfmmwjmmnmmcmtmm\n1-2 f: lffwd\n3-7 f: vzltffffqfff\n1-4 g: grgr\n14-15 r: rrrrvjrrrgrrrrtrrr\n13-14 f: bfswfmfffvffpdffkrhb\n1-11 s: ssssssssssss\n13-14 v: vvvvvvvvvvvvvf\n1-2 g: ggggg\n6-13 m: vrzwmncmmhmmc\n7-15 h: vhhhwghkwhclrhh\n3-5 w: ccwwwxwcnf\n4-5 j: jmcvjz\n3-4 b: bbbw\n14-15 b: bbbbbbbbbbbbbbd\n14-16 w: dzwwwwwwwwwwwwwxkw\n2-4 n: jnxfznsn\n6-7 c: cccccgv\n2-4 w: xwbbqwlswh\n7-10 f: sbbxcnffpgfmfffgtm\n2-3 z: gccczkzt\n4-5 k: wtskbwk\n9-10 k: kkkkkkkkktk\n9-12 l: llllllllwjllllll\n1-5 c: ccccpr\n1-5 s: ssssh\n4-7 r: dtnrmcfsc\n8-10 p: pppztpllppp\n3-11 l: lljllljllfl\n4-13 l: lllvlbllphlldn\n5-6 h: fhhpmhm\n4-6 x: xxxhtx\n5-8 j: mrthqkpj\n6-11 k: kkkmqvfnjkwkxpxkskc\n10-11 j: jjjjjjcjjljjjjjj\n2-6 x: ctmtskbcxbn\n12-15 p: pppppdppnpppppwk\n2-12 q: qqhzqvjnmcmrfkphbrx\n2-3 l: llhxljlcxdvrwppdm\n7-15 j: xbjjhssjjjjpjjjjj\n1-4 k: kkkn\n2-7 v: mpvwbmvvhxjpv\n4-6 k: kkcgwkrlqbbpgqmlk\n1-2 v: khpskvgxvdpw\n7-9 p: ppppppgpt\n11-12 h: hhhhvhhhkhhsh\n4-5 n: dnngnlpn\n3-5 k: kmwkkk\n7-12 s: ssssdsbjbssstssss\n6-12 c: cdcccmcccccc\n3-4 g: gglsggggggggggggggg\n5-14 p: ptrnpppbwwjmckbwpmw\n1-4 t: tbfdtrtbtdjnk\n3-7 k: jwkkzlftt\n1-4 z: jzzzl\n1-5 t: jtttpdhttttgtnt\n7-13 n: nnnnnnvnnnnnk\n5-11 f: ffdfftffffkfdf\n8-14 q: qvcqfqqtkqqjqgqjqn\n10-13 b: qbbbbbbbjqbbbb\n8-9 l: blsljwlwl\n11-14 g: ggxgfrbgggqgvgznzg\n10-17 t: rttttctxrttttttttt\n4-5 j: jjjjs\n4-5 q: qkqzglqgnfqpp\n10-11 b: ddbbbbbbbbbbbh\n7-11 s: wsssssssssgs\n3-14 l: lvxlfjkqllfmzs\n10-11 z: zmznzzkzzhvlzcz\n1-4 n: nnnj\n7-11 f: dnffqxxgznknfcffff\n5-17 k: fkjkdkktkpwwkxkkk\n2-6 q: qtqqqq\n5-10 n: nhxwndndnn\n14-18 x: xxxxxxxxxxxxxrxxxlx\n9-16 n: hnnfnvnncdnrxnnf\n6-7 l: lllxllvlh\n5-6 d: hdddxl\n11-15 g: ggggggqgggcgfgcg\n15-16 l: lllllllllllllllpl\n2-7 q: nmtqsfqqlkxq\n4-5 b: bbcbsnb\n9-11 f: ffffffffffn\n2-15 p: qrnpxnpcpppqrppp\n4-5 s: sqszspvggsds\n5-7 m: mmzsmmcm\n9-17 q: qqqqqqqqgqqqqqqqc\n5-14 w: swwgghwwwsvcjqqjw\n7-9 b: ptnhdbzfbbjsjkm\n4-9 h: bhhtcqnhh\n9-10 t: zcttwtttggttn\n3-4 j: cjzjj\n14-18 h: hbzhgcjhhhhzkhhpwb\n1-5 b: bbbbw\n9-10 k: kkkkkkkkkkk\n5-7 m: mmmmcmsm\n2-6 x: jxxdxlxx\n6-10 v: kxvvdvvrxdvk\n2-12 z: zllzlnzfztlszzzct\n3-4 d: dddd\n2-10 h: hhhhhhhfhddhhbg\n1-6 x: xdgfgxn\n4-5 h: shkhrhhdfhh\n1-8 f: zffbfffffff\n4-10 p: cskpcpjmtpcftsblszpv\n4-17 p: pppbppppppppppppppp\n8-9 j: jjjjjjjvj\n4-5 p: pdppppp\n6-12 k: kkkkkzkkkkkkk\n5-15 n: nqnnklnnnnnnnnnnnnn\n3-7 g: qhggbgdm\n2-3 r: cjprr\n9-10 q: vkgcjcnrzqgq\n5-6 m: mmmmmmm\n4-15 f: fffcffffffffffk\n5-13 f: hfffntlfffpfffffff\n8-9 p: dpjppppppp\n5-6 j: njsrjj\n2-3 k: kkkzzb\n7-8 g: cggbpgkgjgmggvvgn\n12-13 w: wwskwwwwwwwwjww\n6-8 w: lwwjwwwcw\n4-5 l: lllll\n6-7 g: sggkggkplxgxprsnggg\n5-6 h: hjrhzszvhhq\n9-17 w: wwwwwwwwwwwwwwwwdw\n16-19 c: ccccrccccccdcctcccn\n3-7 t: ftngtnntttttttzftttt\n11-13 q: qqqwqqqqqvqqg\n4-12 s: ssssdjssssszxs\n1-14 c: chsckjmxwghqvvckvf\n9-18 s: sssssssslssssssbsss\n7-8 n: nnnnnnnnn\n7-9 c: ccccccwcc\n14-20 n: nnwnnznnnnnrnnrnnnnm\n13-14 h: hzhhhhbhhhhhhshhh\n5-7 x: xxkzxmqxszf\n7-13 l: lllllxllftllhzl\n1-3 f: fftff\n11-12 t: tpktttttcvtt\n6-8 r: frjzrpbrrrqqr\n7-12 w: nwwlpvwwdwbzbldnd\n13-14 r: rrrrrrkxrnrrnrrrr\n13-17 b: bfbbbzlbbbbbbbbbb\n4-7 g: ggggghrpgrg\n1-3 v: jvvv\n8-9 l: llllllllll\n13-15 r: rrrrrjrrrrrrrrg\n4-8 b: bbbbkbprh\n6-14 k: klkzkkkkqkkkkkkkm\n9-10 l: lllllgllsl\n4-6 q: qqqtqqqq\n6-8 l: zkfbhwpq\n10-15 v: rvvjvvvvvvvvvdvvsvd\n2-12 b: bwbbvsrbrbzbb\n9-10 m: mmmmmmmmmmb\n4-8 r: xrzrzmrrzrrzr\n1-10 c: cccccccxbbhcccrzc\n5-6 q: bqqvkqqs\n3-13 t: ttttbtttttttlt\n2-4 z: zfzz\n1-8 v: qvvvvvvv\n3-7 g: ggwgggg\n5-6 d: ddddddd\n10-13 c: ccqccccccscfccjcc\n2-7 h: qghlfphcr\n9-11 w: rwwwwpwwqwx\n2-10 h: shftplxhvplqr\n1-3 n: qsnndvnzczn\n17-18 b: bbbbblbbbbbbbbbbbm\n4-14 k: hjkkqgkdvfjflkkv\n10-13 k: wbtkgkjrdlsxkw\n2-4 h: xhqqmhhvkmbpqqsnrn\n9-20 z: zzzzzrzzzzqzzzzszszd\n4-5 x: mxxxrx\n13-15 w: wwwwtwwnzwqwvwwwwwwj\n10-16 n: tcnlnsqnndnwrnjnnn\n3-5 c: cbccrc\n8-11 z: zzzzzzzzzzhzzz\n5-7 f: ffffffff\n7-8 q: qqqqqqfq\n6-7 v: vvvzvlvv\n1-2 t: zttt\n12-17 b: bbvbbbtjbbbbbwbbd\n10-11 h: hhhhhhhhkkh\n17-18 g: gggggggggggvgggggg\n4-7 t: tttmhttttgd\n10-12 w: wfwwwdwwwwwwldc\n4-6 d: dsdddj\n9-10 v: lkvvvvvflv\n6-8 j: jvjjbjjpvjjjj\n8-12 h: hhhthwhhhhhh\n16-19 x: xxxxxxxxxxxxxxxbxkq\n1-13 c: ccgnccvzczzcphcp\n4-11 n: pnnnnbkmnbqqn\n13-16 x: xwxxxxxxxxxxxxxx\n5-6 m: qffrmmclzbtfqmxb\n4-9 k: kfwzkvjzqk\n11-14 j: jjdjjjgjtzjjjjjjxjjm\n1-4 b: bbbpdqbdbbgdhcbv\n8-9 w: hvnwwwkwmbwpwtbpw\n3-4 h: phhmh\n11-12 j: jjjjjjjjjjjj\n5-15 v: vvvvlvvvvvvlsvv\n5-9 q: gqqqzqhjpqqvqj\n12-13 x: xxxxxxxxxxxxl\n7-9 x: gtxxxxxxxw\n13-14 n: nnnnnnnnngnnmnn\n2-6 s: ssssssss\n1-2 b: bhbbb\n6-7 t: tttttttt\n1-9 k: kvkcjwzhl\n8-11 b: bbbbbwbbrbk\n15-18 p: pppppppppppppppppp\n1-6 f: wfnftfvff\n2-4 d: tdddddpdw\n3-16 s: ssbsssssssssssspz\n15-17 p: gpxpppppppppppppg\n4-5 w: wwwww\n9-12 t: htttttttzlgtrtt\n3-5 c: fcncc\n1-3 g: gggn\n6-12 c: kcccccxncwcqc\n4-6 v: vmpvqv\n1-8 x: sxxxxxxdlx\n16-17 d: dddddddddhddgjdxdddd\n5-9 s: sbkrswskxsskskqlcp\n5-8 h: hfhhhhhfgd\n2-5 t: tdttt\n10-11 k: kkkfkkkkkkp\n10-11 h: hhhqhmhhhjl\n9-15 r: rmrrrrkncrxrrrp\n5-7 k: kkcqxkkkj\n14-16 g: gvgxgkgpgrxsgwxw\n9-13 g: qbwsggggzggggp\n2-7 t: jdvbklttd\n1-3 r: trrfxqff\n7-8 t: ttxtttvt\n16-17 b: bbbbbbbbbbbbbbbbb\n5-14 r: rgrrlrrrrrrrrrgr\n10-11 x: qbfdztwxpxdxj\n2-3 q: brqq\n5-15 t: ttttftvlttttttgttt\n3-6 v: hvhthfw\n4-6 n: nngdnkxzc\n2-7 l: bkwmlrl\n2-8 q: xtqqqqhdwq\n13-15 f: fffffffbfvffffflf\n6-9 c: dqvwbscpccccrcccchws\n4-5 l: gvjltw\n3-6 q: qgnpjqgbl\n1-3 t: tsgt\n6-8 p: ppjppphfpj\n1-4 d: xddd\n12-13 m: mmmmmmmmmmmmvm\n7-15 b: bvnbcrbvtpbqbbbbl\n13-15 x: kxxxhxxvxxxtngxxxxxx\n4-5 p: pppvvtncp\n11-14 p: pppvpppppspdpppx\n2-3 p: wppp\n8-9 f: fdmcfffxnnff\n9-11 q: qqqqqqqqrqn\n13-15 b: bbbbbbbbbbbbbbv\n6-12 r: rmvrrlrrjrrrrrsrrr\n1-2 l: qllllpkllhlllbw\n9-12 w: pbvwkbfwdngw\n3-13 k: kfhdkkkknkkkkkkhl\n8-18 b: vbbnbqzpjwbbdbbcglsr\n5-17 q: qqqqqqqqqqqqqqqqqqh\n4-8 c: cccwbcncdcc\n8-9 n: kfsbvbnntcn\n2-5 f: vfvffrq\n17-18 j: jjjjjjjjjjjdjjjjtp\n5-7 g: ggjhgcbg\n14-15 q: qqqqqqqqqqqqqqp\n2-3 d: dzjd\n6-18 t: ntntttqtmttttktttqtt\n2-6 r: rmwrwppprc\n9-14 z: zlzzzzklsgzvfz\n7-8 k: vkfglktkkkkk\n3-8 d: mdkdxxbdqddrjwvc\n13-17 c: ccfccccccccczccvcc\n4-5 k: kkknrr\n12-15 s: vbssssssssnfsssss\n7-18 m: mmpmmtwvkmmcnscmdmj\n13-16 v: dvvvrvdvvvvvjvvvpjvc\n3-4 t: tjtkndtbwh\n4-6 k: kkbkpz\n8-13 k: kckzkktnfzmdkjkkb\n8-13 j: jsnntcjcjjqzjj\n11-13 b: bgwbbbpbbbqbbb\n3-6 r: zqrrrjsj\n2-7 k: zkklkttkxk\n3-5 p: gjphpsppntgp\n17-18 l: llllllllllllllljlll\n6-11 r: srxfbrswrnlfp\n1-8 q: nsbqchsqqx\n5-7 r: rvrrkjcrjmr\n5-12 r: rrrrrsrrrrrvrrrr\n3-7 k: kkdkkkkxk\n11-16 p: rppppppppcppdpmbp\n16-17 p: pppppppspppppppjpp\n4-15 t: zsrtxgfdtrgxhmcst\n5-10 k: gkkzkfkkkknk\n13-17 w: wgwcwwwwwfcwmwwhpww\n7-8 l: llllllvw\n2-7 g: ckdqngmzhghcvlcn\n14-15 q: qqqqqqqqqqqqcbqq\n3-6 m: mfmcmd\n3-6 b: bbbbbm\n1-8 n: nmwnnsnmznnnmrn\n7-11 n: lnnnmrnqnntxnqcn\n3-12 h: tghhthwtjjhhfbhdpnh\n2-4 d: jsddd\n8-9 w: wwzwwwbwzwc\n10-12 r: rrrrrrrrrrrrr\n1-2 b: xbbqkfp\n3-6 p: dfppptptv\n8-12 z: jzzzdbzhwzzmklzkvmz\n1-4 n: nqnprp\n2-11 b: jbshxxbxczhscksvhctm\n5-6 b: bmhbcvgjfbc\n3-7 m: qlfmmdm\n8-9 v: vvsvsvvfv\n7-10 l: llllldllln\n9-10 c: ccccfgcctcc\n7-17 j: vvjgnjvtjjjvvjjjtj\n5-6 s: ssssvs\n3-11 q: qqqbqqtqqqzqqgj\n4-5 k: rkjskkqk\n4-16 k: kkkkkkkkkkkkkkkkk\n2-4 h: shch\n5-7 n: kznnmgnnnn\n3-8 r: xhrrrrrg\n4-6 m: bmmmmgqmnn\n3-14 t: tttktwfmthkttdq\n4-5 w: wwwvww\n1-2 k: kkrkmkkw\n6-11 h: hhhhhhhhhhsh\n1-9 h: fhhhhhhhbhhhhhhhh\n7-9 b: frkqtcwstzqbdd\n2-3 p: xxpvrmcmppbw\n12-14 r: rrrrrrrprrrrrq\n5-7 l: llcqllc\n10-13 x: bpwxpmrrmxbjrv\n8-16 w: wfcnvwvwlvtnwwfg\n1-5 k: kglsfqshkpbs\n1-8 w: mwwwwwwww\n4-7 d: ddgldhddvd\n5-7 v: qvjvxvvvvvvv\n1-2 s: ssns\n6-9 m: prmmrmbmqmftmmb\n13-14 p: ppppphpppppppq\n8-10 k: kkzkkklkkkck\n1-3 p: pnpswpmsfk\n15-19 r: trqmkzfvlrkrrgsjhrr\n15-17 c: ccccccrxcccscccczcxc\n2-6 b: rbbbbdpprw\n7-14 l: lllzlsklllllll\n12-14 l: lnllhlllplltllrplll\n3-5 w: wlwwbr\n5-9 h: hhcvppjrhhhhhqc\n2-3 w: wwwzt\n2-5 l: lnrlll\n1-4 v: vvvl\n5-6 g: mkngggzgxkggkldg\n6-7 m: mmmmmlm\n8-10 k: cnkkzkklwkvkkk\n7-13 g: ngqggglgggggggg\n8-10 m: mmmkmmmpmm\n8-12 f: ffffpffffffz\n12-15 f: ffffffxffffcffff\n9-14 w: wtrcwthdwxnwwfpwbj\n4-7 l: xfllfnlxlgbll\n6-8 c: tcdppkscmc\n4-10 s: qssssspsskss\n2-4 t: wgpdqdcdtd\n3-4 s: shsbksdj\n3-4 n: nnmn\n2-6 x: dxxxxjxr\n1-3 b: qbbptm\n9-11 h: hhhvhhhhvhhxhhhrr\n3-5 x: xxxxw\n2-4 l: kgrll\n3-5 m: mmhmmmm\n7-9 j: fjkjbjjjfjhcj\n3-4 n: ngfnbzln\n9-10 x: xxxxxxxxxz\n11-12 c: ccccccccccjcg\n9-14 j: swnjssjjjjjmjnj\n2-3 n: nrcfrf\n11-12 g: gggggggggdzxg\n4-5 v: vvbvnq\n8-9 s: ssssssssx\n6-7 l: zlznxjlb\n14-16 x: xxxxxxxxxxxxxdxx\n6-8 x: xlxjxxlxxmxx\n2-7 r: prgnrrrqrpcr\n15-16 m: mmmmdmmmmmmbmmhmmm\n14-15 w: wwwwhwwwwwwwwwwwww\n5-9 k: kkkkkkkkkk\n8-9 c: rccccqcchc\n6-12 v: pvvvvvzlmvltkh\n10-11 h: hbhslhhhhlhhhh\n1-7 s: ssjswbj\n5-12 q: qslqzqjqxxqpqzvjnq\n4-11 z: znmzbhdgftf\n5-7 g: sgtwggggghgglfgt\n2-4 d: kdddfdqk\n11-13 m: mjmmmmmmmmmmt\n12-14 k: nskkkkxkkskwkk\n6-11 j: vhpjqjhrjjjg\n7-14 l: lllllllllllllm\n2-15 g: bcqpgsvgkfgsmrgvpgng\n2-9 w: grwwwwcbwwwwdxprt\n3-18 w: wwkwwxvwwwhnwwbhzw\n5-13 g: ggggggggggggdgggggg\n10-12 h: hhhrhhhhcldh\n2-5 k: xhpnk\n11-14 f: ffffffffffffff\n4-8 x: wxxjxxmx\n16-17 d: dddddddddddmdddvd\n3-14 l: lllllllllllllpl\n1-2 w: rfzvvcww\n5-7 v: nwvxvxkvvxctvh\n1-3 k: kxghkkb\n17-18 j: jjjjjjjjjjjjjjjjht\n2-5 t: tchptl\n9-15 j: jjpjjjjjjjjjjjjq\n1-7 r: rrrrrrxrrrj\n13-15 h: hhhhhhhhhhhhhhqhhhhh\n2-4 v: xdwvk\n3-6 w: rwmwkwgcxzjtwww\n2-3 w: nwkwtkw\n3-8 d: jddhdddlqd\n10-12 m: mmmmrmmmmvmmfmm\n2-7 l: llllllkllln\n9-15 x: xxxxxxxxpmxxpxxxwx\n4-8 n: nqnqnnnzn\n7-12 d: zddpjdddgddgxddwdddw\n7-8 b: bbwbbbrl\n6-13 p: zppkppprpvnbzpp\n2-8 w: cmnrmkzwpxrxwhfsd\n5-12 x: xvrxpbhlgfxx\n11-14 j: jjjjjjnjjjjjjc\n7-10 r: rdsrhrtrrrlrfncrb\n15-16 x: xxxxxxxxxxxxxxcx\n7-9 h: fhhhhrhhhh\n4-18 v: vvqjvvvvkvvqvkvvvvvv\n1-7 x: xxcfjxkxtxx\n10-12 c: ccccmccwclcjc\n5-6 l: llllkl\n1-5 h: hhhhn\n1-5 p: bppppp\n17-18 j: jjjjjjjjjjjjjjjjjj\n9-16 k: fkkkkkkkkkpkkkkhkkkn\n13-14 t: ttttttdttttttsttvt\n7-16 m: mmmmmmlmmmmgmmcmm\n14-15 d: ddddrdddvddddhdd\n2-4 x: xbxxtqxd\n5-16 f: fffffffffffffbfff\n12-16 w: wwwwjwwwwwwwmwwn\n5-7 c: cclccfvzccccmb\n7-16 m: mmmmmjmmmmcmnswj\n4-10 x: qpxrcxpxxsxx\n15-20 d: dddjdddddddddddddddx\n2-3 x: cxmpxs\n2-5 t: tlhztr\n9-12 x: xcdxhqhxpxgx\n1-11 f: zxnfsfjftfffjfzswdf\n1-3 j: pjmk\n15-16 h: kpcschhvfzphpnhvhhb\n16-19 s: ssssssssssssssssssf\n5-7 p: kppdcppp\n4-8 f: fffjfrfzfffmffjf\n2-5 q: zqqkvkxvzm\n11-12 w: wwwwwwwwwwww\n9-10 f: rffffffffzf\n12-13 p: pppppcpdppfrp\n7-9 n: nntnnnsnnnnnnn\n9-12 k: kkjkkkkkkkkkk\n3-6 m: mmmmfmm\n3-4 v: vqvw\n9-12 b: bbvbbbbphbbn\n4-6 b: mbbbplbvgbdgb\n6-8 h: hhfscsmznfccc\n1-8 h: khtshhkhwfc\n1-7 x: vxxxxxx\n5-8 l: llllrlllll\n9-10 l: llllllllll\n5-9 j: jqscttkjjsjjjnd\n7-14 p: pnlprpppjndpphppppp\n8-11 j: jxmjjjjjjjsjfmsq\n8-12 h: hhhhxgmhhwpfh\n4-13 h: hhhnqzhfmshhhhshhh\n3-4 q: qwtq\n7-8 c: fcczpcgqnccsc\n3-4 w: gwcwrl\n6-10 f: ffkfffffnffqcf\n7-8 f: fwvfmvffwfff\n5-9 l: llllxllln\n10-11 r: bjrdxrrbrrbrprrrcrd\n4-5 g: gggggtgw\n8-9 r: rghdrfrrg\n9-10 v: vvvvvvvvvv\n4-5 d: ddddlrzhdgw\n2-4 x: zxlx\n5-6 j: xjjjwj\n4-9 x: xlkxnmjnxtzmxxdqdxk\n13-15 f: ffwzfffpspfgxffffffp\n4-6 v: vvjxmrvrf\n4-16 b: ghzbqjqbhtjbbmkvkrb\n7-12 q: qqqqqqkqqqqqqqq\n4-15 z: zcmzzzlzzrlkzzzg\n13-14 x: xxxxxxxxxxvxxq\n2-5 h: hhgfhjxbpwmthtjsjhhl\n5-6 x: xxxltx\n4-16 v: vvvvvvvvvvmvvvtnvv\n9-10 s: gsvhnlspts\n2-4 k: kkkr\n4-6 j: jjnjjjtwnmjv\n4-6 n: nnnsnn\n3-4 d: fdjf\n6-8 r: rrrrqvdrrrsrr\n6-10 w: wnqvwqwqwwb\n9-10 k: kkskckkkkk\n4-18 d: dgdjjdddsddwdbhddm\n8-9 b: bbxbbdbkbbdb\n15-16 l: rvphlrrpmwfqkmcl\n3-4 m: mbmtxznmg\n4-6 c: cclccxcccccc\n8-9 p: ppppxbpqpppjp\n10-11 c: ccccccccccc\n4-13 c: cdnccclcszcwfcj\n5-8 l: ldllkllj\n2-10 k: dkbgkwgkzwwtkkc\n3-4 p: qrpcpdsp\n5-8 m: mmmmcmms\n2-5 c: czcpcc\n4-9 k: bkkkdkkkgk\n14-15 l: lllhllvllxbmgbl\n14-16 j: jjkzjhjjjjjjfjjwjpjw\n13-15 w: wwwwwwwwwvwwwwswwwww\n6-7 q: qqqqztv\n9-16 m: mmmfmmmmmmmmmmmmmmmm\n3-4 m: hmvnmjz\n12-13 g: gqggggggggglgggg\n3-7 h: hhqhhgkhshh\n7-9 h: jvptznhrrbhffcdp\n4-6 t: tttttttttt\n2-8 f: rfhbmftfxnxllkpqh\n6-14 z: zwzzzrzzzzzbzzzzgzz\n5-8 t: jkmddtqcqttrtzk\n2-5 p: ppprqsgpzng\n4-5 v: vvvlvfrjxh\n14-15 c: ccccccccdcccccvc\n5-6 j: jjjjjs\n1-2 g: ggdnzh\n8-15 p: ppfpgpxdvphdcpppp\n2-8 h: hmhhhhhthhhhh\n3-4 h: hhmn\n7-9 b: bbfdbjbqb\n3-6 d: ldfpftkxwqddbcdqd\n2-7 j: wstjkjjdjtpwwwxbg\n5-8 p: vmdppswpppzpqxrdt\n4-8 p: ppsgprrcpskp\n10-11 r: rbrzcfrrvmrtrrrrrcr\n5-9 j: jjmjjjjjx\n10-16 s: sssssssssssssssls\n4-7 p: ntpfphp\n1-2 l: glll\n9-11 x: xtxxxxxxrxxx\n7-8 d: ddnddxrrd\n1-3 x: jxdxfsxxtzvxxwxx\n8-9 s: ssssssncsl\n5-10 s: tsmspsxszsfftcs\n2-5 b: bbqbqb\n3-5 w: dnxljjsclwg\n11-12 n: cmlhcbxxnnmhn\n11-14 g: ggtgggdkhggggbggrg\n2-3 m: jvvm\n1-12 g: gggggggggggxggzggf\n12-13 p: mpppjpplppvpqpcpvp\n5-16 w: swtwhmxzwbwxwwwwwww\n5-6 m: mmmmzmm\n8-9 v: vvvvvvvvg\n3-4 w: lhprqpcwf\n9-14 j: jjjjjjjjjjjjjfj\n2-7 s: dssgssswp\n10-11 b: ldfwbqqhgbbpbbrcs\n11-12 f: bffvnfgfffqfqdgffzh\n2-8 p: hfcppzgp\n3-10 r: hfrcwtsrcmkrn\n4-5 j: jjxsjjj\n2-4 g: gggjg\n4-5 z: zzzcv\n2-6 w: wwncqxp\n5-6 m: lmmrjmmmmqm\n1-19 d: lddzdtdkdvddddddddd\n16-17 q: qzqtqqqqqqqqqqqpqqjs\n6-7 c: cfccxxc\n3-5 c: qgnccb\n6-11 s: cqqssmxwznspv\n11-19 s: sssssfssswgdsshssszs\n15-17 x: xxxkxxxxxxxpxwxxhx\n2-9 b: zbggdxbwbpv\n4-10 z: zzrjzzdpxzz\n12-13 m: mmmmmmmmmmmlm\n4-8 s: ssswssss\n9-10 b: xdbhbbpbbhkbbblvb\n11-12 t: ttttttttttlj\n5-6 l: llsfjll\n14-16 x: xxxxxxxxxzxzxxxzxx\n2-10 h: hsdtlhhpwhhh\n3-5 s: shgbh\n4-5 j: jjjjps\n13-14 d: ddddddkskddddddlmwdd\n5-10 w: hbwrwqwzwlgz\n3-6 b: bmbbxjbw\n11-12 f: fffffhfffffr\n10-17 l: llllllcllllllllll\n3-7 z: zbmzzzxzd\n4-5 l: dlbld\n2-5 r: mghzsbjrsqflrd\n6-8 q: qqqlqqqtql\n8-9 s: sssspsstss\n5-6 w: wwwwwm\n2-4 m: mxmm\n1-3 s: prrzbwfslzpdks\n10-13 v: vgvvzvvvvvnvvq\n10-11 w: wjwrwwwwwww\n14-15 b: bbbbbbbbbbbbbtk\n2-7 n: nnnncnn\n5-10 w: wwzwrwwwwwwww\n8-14 f: fffffwfznrvflf\n3-4 l: rllr\n2-4 x: phxkxxxx\n2-13 w: vwtlwgwmwwwjd\n3-8 v: xdszpbvvmvxbhcvlb\n1-6 x: xtkkpzdwtxx\n7-10 x: xxxxxxxxxkgxxxxxx\n5-6 q: qqvqqd\n18-19 m: mmmmmmmmmmmdmmmmmmz\n2-5 s: bbsts\n9-12 l: vlllllllcllll\n1-7 x: xxnvvxxpx\n2-3 f: dfbf\n4-6 v: vcbvrm\n1-3 k: ddkk\n1-6 q: qklqqqvs\n7-8 f: ffffwfhfmff\n3-6 b: wbbcsqhnpb\n5-6 x: xxxxxs\n3-10 g: gnsgrhrgsgwrlgcwgh\n4-5 t: tttht\n4-5 l: lllqllllllllllllll\n14-15 z: zzszzzzztzszzzzczzz\n5-7 j: cqvjxzjslrdjnjjq\n11-12 d: dpgddddzdddrdddq\n14-16 w: wwtwwzwwwwwwwkww\n9-13 p: pfppppfrqpppc\n3-4 x: gpklxz\n3-6 p: pxppznbpnp\n4-16 c: bdvcfddsvccqvcwcmgc\n1-2 s: dsgbs\n8-10 w: wwwwwwwwwf\n10-18 b: bbcbbbbbbfbbbbkbbb\n3-7 w: kzwcwtwjhb\n16-17 d: dddddddddddddddmh\n2-4 t: ttttx\n2-9 f: fwnfgfffftbf\n12-14 t: rttttztwtttwtttttc\n6-8 d: bfdjhhdmmffbdd\n16-18 f: lcnvfgvfggvtttnxjf\n2-8 l: llnsqgqcqwjs\n4-5 f: ffmdwsrfp\n3-4 q: fbkqkrlsqvlqmxqvv\n14-16 w: wwwwwwwwwwjwwdwhww\n1-4 t: tttnt\n3-5 s: vmbwsvqdssgrrzbj\n12-13 q: qqqqvqqqqqqqjqqh\n5-17 n: nnvnnnnnnvnnnnnmg\n2-4 l: fllll\n3-5 t: tqtvt\n6-8 r: rrrrrxrs\n1-4 z: zzzz\n1-2 k: gkbd\n1-5 r: rrrrg\n2-11 l: kwwjlllgwhlnlsqtrq\n13-15 h: hhhhhhhhhhhhghh\n1-3 h: hwhhwhxhh\n5-10 c: cmrgcfccccccj\n17-19 c: cnccccccpcccccccmcz\n1-8 v: hvdvgbmvxnvsvgkvvv\n10-11 k: kkkkkkkkkkj\n3-6 d: rjdjdggfdcxpldr\n3-14 g: ggjgblgggggglggggg\n18-20 z: dwzwsrqnxmzhswfgjzsz\n1-10 s: qshsqssmtsssbgsg\n7-17 l: hmllllqfkllllpllnlbw\n6-7 w: wgcwwnwwwx\n1-15 k: kvskkbkkkhkkktkkkk\n6-7 q: qnqmhnq\n2-6 n: nnnkxwjhfmtnfnnmg\n4-7 r: rwrfrrr\n13-14 n: nnfnnnnnnnnnpnnn\n3-5 h: hqktph\n8-12 f: fmffhjmfkxfrvcfrbr\n12-13 j: jjjjjjjjjjjljj\n2-9 l: glcwqflmllstjhpsp\n4-6 r: vrwrrxrm\n5-6 n: nnnnnsn\n14-18 l: lhllclllzsllvvllzlll\n12-13 f: fvzfffnfpxffm\n7-8 c: ccrcccfpwcg\n1-2 g: tggg\n9-13 x: bdxxxnhxmxxxxxxnd\n14-15 z: zpzgzzzvzdqzzzt\n13-14 b: bbrbbbbbbbcbmbbbkc\n1-2 f: fjndbffq\n8-11 h: zhwhlrhhhhlnkhhh\n3-8 v: vqvwvclvcplmvctjv\n6-7 g: ggxgglggngdg\n4-6 h: vslffhhwcbnhh\n6-9 c: rztcpdvcpc\n7-9 t: qtttttktttlk\n17-19 d: ddgdddxdddtmzddgdtwf\n16-18 q: qqlqqqqbqqqqqqqqqqq\n3-7 q: lwxcqqtjqqbs\n1-10 w: txwcwxqxzw\n1-2 v: qvvvvvvvv\n6-11 c: ccccchcccccctccc\n6-11 d: dddnddklzdddd\n3-12 p: pkqrppgcmptp\n13-15 r: rrrrrrrrrrrrrrt\n3-4 k: kkkmjkg\n6-11 l: xlllllsllmqncvc\n6-17 g: vgggngrgggqggggggg\n3-4 b: bbqb\n3-9 q: qfqdqlpxgtbqqjmjccqt\n4-6 f: cbfmfc\n2-11 d: gddtddtdddkpdfcdwm\n2-11 q: qmpdmfhkqlql\n8-9 v: vvvvgvvhv\n2-3 h: hxhhhp\n9-10 n: nnnnnnnnhthn\n16-17 q: qqqqhqqqqqqqqqqsk\n6-12 v: vvwvjbjdvxvzvvv\n10-12 v: vxvvvvvvvvvcj\n3-4 h: znhfvctppgjtqhhhl\n14-19 g: ggggggggggggggggggcg\n7-17 p: rpnfbxppcppprbppkds\n4-11 c: ccccczcccclc\n2-3 d: dvdl\n18-20 t: twtthtdgtdsvxttcfltg\n11-12 q: qqqqqcqhqmqgqqqs\n10-11 c: ccccvccccgc\n7-8 x: xxxxxxhxxxxxxx\n8-10 r: rcnrrrwfcz\n1-3 b: pjbk\n6-7 q: xqqqcbqjlq\n11-12 z: ndzzzzdzzzrzz\n2-9 g: gsgbmggnggggkfbcgggg\n11-13 k: qbfnkxbkkkhkkkkkkpkk\n4-7 d: dddvddkd\n4-10 c: jccdcpfchcszccccxfc\n8-9 z: xvzvnzzzfdzlzr\n2-5 n: zqnqnpnhwnnv\n11-12 g: ggkgggggkgxgg\n3-4 v: vvvnvtgvpr\n2-4 t: gtmtxkbqvj\n4-5 m: mmmmtmmm\n10-13 g: ggggfgsgpcgggvggg\n4-5 j: jjjjj\n12-15 v: vgvvzvvvvvvnjzvv\n16-18 r: rrrrrrrrrrrrrrrrrrr\n11-13 r: rrrrrrrrrgfrr\n12-13 s: tsssssssbssdss\n12-14 m: xmmmmmmmxmlpmmmmm\n6-13 w: wkwmlwwhcgwmgwjwgwww\n9-13 b: bbbbbbbbbbbkbbbg\n4-5 v: vttvlvs\n7-10 b: mbfbbbrffbzbbxlx\n4-6 t: ttcfxtns\n8-12 r: dcrzvdpprbmr\n2-5 j: cjrxjjfz\n9-12 b: bbbtwbnkpbrrkbbbqlb\n3-4 k: kkdhk\n4-5 v: vjlvkjxvv\n3-4 x: xxxrs\n5-7 h: hhrdhhhhh\n5-6 v: vfvvhfhc\n3-8 w: wwjwwwwwww\n9-11 g: gggngtgggghg\n1-4 q: hcqq\n11-18 w: wwwwjwwwrwblwwwwwwz\n2-10 h: chrvlwffjl\n13-14 w: wwwwwmwfwwkwwhww\n5-7 z: zzbzdzqz\n4-13 v: dvxclvvxcnmcxvsp\n7-11 c: fccwdcccsfbcc\n12-16 f: pfsfsgpcfffffffvffpf\n1-4 n: gnnnn\n6-8 j: jjjjjjhbd\n3-4 k: ktkp\n6-8 m: rmmgmjsmmm\n7-8 j: jjjjjcjjj\n15-18 r: rrrrrrrrrrrrrkgrrr\n1-2 z: zzmwrlmtwsxdbzss\n2-6 j: bpwbjqr\n3-14 q: njswqhttbqfqqzqhtq\n12-17 d: ddddddddzfdvqddmt\n2-15 j: jfxjjjzjjtbcxtj\n7-10 w: dpzwqflvdx\n2-8 b: bjbbczbbwbbbgmdvllb\n7-11 b: bbbbbbmbbbb\n7-8 f: smbzfnxgvm\n6-10 n: nnmnncndtnnn\n13-15 c: scnqcccccccqpgccqccc\n3-4 w: wwwv\n3-7 h: hhhrxbx\n8-12 n: nnnnpndnncts\n1-4 g: kgggfggg\n4-10 m: fbttqmzmmmxtmtm\n2-4 d: dmclzdvdd\n10-14 s: sfgtsjfsssssstbs\n1-3 p: ptpzpqlhprvxhrgvvgv\n4-7 n: nnnvnnnn\n5-6 c: plnccqrxtxcwzkccm\n5-10 j: jjjjjjjjjr\n7-9 b: bblbbfhbb\n7-10 t: wtttthtttthvt\n2-7 m: mrmfmjgcpmmmbwx\n9-16 k: kzklkwkkkhnrkkkkk\n9-13 l: qlllljlllllllllkk\n2-3 c: ctcc\n5-7 b: bbbbbbxb\n3-6 w: wtqrwwgkv\n5-7 x: xxxwxxp\n2-3 z: zzxcdn\n13-17 c: ccccccrfcccccnccx\n3-7 w: whfdzvwpg\n15-19 t: ttttttttttttnthtttct\n10-11 k: kkkkkkkkvxkkk\n4-9 g: gggggggglg\n3-5 j: ljnjj\n4-8 l: nttlhlldhsslcqzqpgdv\n7-8 r: tpxhclrrsdnmwcrgf\n8-11 w: lwwwwwwwwjw\n8-9 w: wppkzcrdmkwgqw\n3-8 z: vhzgnkzgkcqrmmvvkx\n7-11 n: nnnnnnnnnnmnn\n1-15 b: bbbftxndbbfbbpb\n6-7 d: wkdddldldcddddd\n3-5 b: bcbfvbd\n9-10 k: kkkkkqkkrk\n8-18 z: vszwbmgzjzgjmhpgcv\n4-9 z: zzzzznczxz\n3-4 d: dhdm\n1-9 q: zqqqqqgcqqq\n2-3 g: zvgnx\n2-4 j: jqjjh\n6-11 j: hcqcnjqjjgj\n5-6 j: jjjjjj\n6-7 d: dwdddkd\n1-5 l: pmplllll\n1-11 t: tttdlttsrpkqtt\n3-9 x: xxnxtxxxx\n7-8 f: fffffffg\n4-6 b: bhbbbrhbb\n6-11 f: xfncmzffrfsf\n10-11 h: pghprtcjjjhshkw\n8-10 n: nnknnnnznnnn\n3-5 q: jqqqk\n5-9 j: rrhfgjfjjjjjbdjnj\n6-8 t: tctjmtttqttt\n9-10 m: wmvmhmmmxddzmmmm\n2-11 c: flcqrnrqmcccs\n2-3 f: qfdx\n11-15 k: kkkkdkkkkkgkkkkkkrq\n9-14 h: hzhhfhhxhhhhhltnh\n"

//dataRepository holds a slice that contains all Data
type dataRepository struct {
	allData []*Entry
}

//Entry is a Data Point that holds the minimal and
//maximal digit plus the letter to check for plus the password to check
type Entry struct {
	min, max         int
	checkLetter, pwd string
}

//insertEntry inserts a new Entry
func (v *dataRepository) insertEntry(entry *Entry) {
	v.allData = append(v.allData, entry)
}

//newEntry creates and returns a new Entry
func newEntry(mi, ma int, check, pw string) *Entry {
	e := &Entry{
		min:         mi,
		max:         ma,
		checkLetter: check,
		pwd:         pw,
	}
	return e
}

//readAllData reads the data of the puzzle input and processes it
//to put it in a slice
func (v *dataRepository) readAllData() {
	//replace all symbols that we don't need so that we get clean fields
	cleanedString := strings.ReplaceAll(puzzleInput, "\n", " ")
	cleanedString = strings.ReplaceAll(cleanedString, "-", " ")
	cleanedString = strings.ReplaceAll(cleanedString, ":", " ")
	separatedString := strings.Fields(cleanedString)
	//every 4 ids we read and insert our data into our dataRepository
	for i := range separatedString {
		if i%4 == 0 {
			min, err := strconv.Atoi(separatedString[i])
			if err != nil {
				println(err)
			}
			max, err := strconv.Atoi(separatedString[i+1])
			if err != nil {
				println(err)
			}
			check := separatedString[i+2]
			pwd := separatedString[i+3]
			v.insertEntry(newEntry(min, max, check, pwd))
		}
	}
}

//checkPwdPolicyTask1 counts the valid letters in the password
func (v *dataRepository) CheckPwdPolicyTask1(min, max int, checkLetter, pw string) bool {
	validLetters := strings.Count(pw, checkLetter)
	if validLetters >= min && validLetters <= max {
		return true
	} else {
		return false
	}
}

//processCheckTask1 processes all data and increments a later returned checkSum
func (v *dataRepository) processCheckTask1() {
	checkSum := 0
	for _, j := range v.allData {
		check := v.CheckPwdPolicyTask1(j.min, j.max, j.checkLetter, j.pwd)
		if check {
			checkSum++
		}
	}
	println(checkSum)
}

//checkPwdPolicyTask2 checks for all three cases that are described in the puzzle
func (v *dataRepository) CheckPwdPolicyTask2(min, max int, checkLetter, pw string) bool {
	checkBool := false
	if checkLetter == string([]rune(pw)[min-1]) && checkLetter != string([]rune(pw)[max-1]) {
		checkBool = true
	}
	if checkLetter != string([]rune(pw)[min-1]) && checkLetter == string([]rune(pw)[max-1]) {
		checkBool = true
	}
	if checkLetter == string([]rune(pw)[min-1]) && checkLetter == string([]rune(pw)[max-1]) {
		checkBool = false
	}
	return checkBool
}

//processCheckTask2 processes all data and increments a later returned checkSum
func (v *dataRepository) processCheckTask2() {
	checkSum := 0
	for _, j := range v.allData {
		check := v.CheckPwdPolicyTask2(j.min, j.max, j.checkLetter, j.pwd)
		if check {
			checkSum++
		}
	}
	println(checkSum)
}

func main() {

	v := &dataRepository{} //create repo
	v.readAllData()        //fill with data
	v.processCheckTask1()  //process for task 1
	v.processCheckTask2()  //process for task 2
}
