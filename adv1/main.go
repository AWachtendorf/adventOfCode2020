package main

import (
	"fmt"
	"strconv"
	"strings"
)

type values struct {
	allValue []int
}

func newValues() *values {
	v := &values{}
	return v
}

const allCopiedValues = "1028\n1987\n1938\n1136\n1503\n1456\n1107\n1535\n1946\n1986\n855\n1587\n1632\n1548\n1384\n1894\n1092\n1876\n1914\n1974\n1662\n1608\n2004\n1464\n1557\n1485\n1267\n1582\n1307\n1903\n1102\n1578\n1421\n1184\n1290\n1786\n1295\n1930\n1131\n1802\n1685\n1735\n1498\n1052\n1688\n990\n1805\n1768\n1922\n1781\n1897\n1545\n1591\n1393\n1186\n149\n1619\n1813\n1708\n1119\n1214\n1705\n1942\n1684\n1460\n1123\n1439\n1672\n1980\n1337\n1731\n1203\n1481\n2009\n1110\n1116\n1443\n1957\n1891\n1595\n1951\n1883\n1733\n1697\n1321\n1689\n1103\n1300\n1262\n1190\n1667\n1843\n1544\n1877\n1718\n1866\n1929\n1169\n1693\n1518\n1375\n1477\n1222\n1791\n1612\n1373\n1253\n1087\n1959\n1970\n1112\n1778\n1412\n1127\n1767\n1091\n1653\n1609\n1810\n1912\n1917\n935\n1499\n1878\n1452\n1935\n1937\n968\n1905\n1077\n1701\n1789\n1506\n1451\n1125\n1686\n1117\n1991\n1215\n1776\n1976\n846\n1923\n1945\n1888\n1193\n1146\n1583\n1315\n1372\n1963\n1491\n1777\n1799\n1363\n1579\n1367\n1863\n1983\n1679\n1944\n1654\n1953\n1297\n530\n1502\n1738\n1934\n1185\n1998\n1764\n1856\n1207\n1181\n1494\n1676\n1900\n1057\n339\n1994\n2006\n1536\n2007\n644\n1173\n1692\n1493\n1756\n1916\n1890\n1908\n1887\n1241\n1447\n1997\n1967\n1098\n1287\n1392\n1932"

func (v *values) addValues() {
	var idasint []int
	idsToDelete := strings.ReplaceAll(allCopiedValues, "\n", " ")
	seperatedIds := strings.Fields(idsToDelete)

	for _, j := range seperatedIds {
		k, err := strconv.Atoi(j)
		if err != nil {
			println(err)
		}
		idasint = append(idasint, k)
		v.allValue = append(v.allValue, k)
	}
}

func (v *values) compareValues(){
	for _, j := range v.allValue {

		for _, i := range v.allValue {
			for _, k := range v.allValue {
				if i+j+k == 2020 {
					println("THIS ACTUALLY WORKS")
					println(fmt.Sprintf("FIRST NUMBER: %v, SECOND NUMBER :%v, THIRD NUMBER: %v ", i, j, k))
					println(fmt.Sprintf("So the solution is: %v", i*j*k))
				}
			}
		}
	}
}

func main() {

	v := newValues()
	v.addValues()
v.compareValues()
}
