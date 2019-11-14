//line parser.rl:1
// -*-go-*-
//
// SDP message parser
package sdp

//line parser.rl:7

//line parser.go:12
var _sdp_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3,
	1, 4, 1, 5, 1, 6, 1, 7,
	1, 8, 1, 9, 1, 10, 1, 11,
	1, 12, 1, 13, 1, 14, 1, 15,
	1, 16, 1, 17, 1, 18, 1, 19,
	1, 20, 1, 21, 1, 22, 1, 23,
	1, 24, 1, 25, 1, 26, 1, 27,
	1, 28, 1, 29, 1, 31, 2, 30,
	0,
}

var _sdp_key_offsets []int16 = []int16{
	0, 0, 1, 2, 4, 5, 6, 7,
	8, 11, 15, 17, 20, 22, 25, 38,
	52, 65, 79, 99, 113, 114, 115, 116,
	119, 122, 123, 130, 131, 144, 158, 160,
	163, 164, 166, 167, 169, 172, 174, 177,
	178, 179, 192, 207, 208, 209, 222, 236,
	238, 242, 255, 268, 281, 296, 297, 298,
	311, 325, 327, 330, 331, 332, 335, 338,
	339, 352, 366, 379, 393, 414, 428, 450,
	452, 453, 475, 497, 511, 518, 528, 530,
	533, 535, 538, 540, 543, 546, 549, 550,
	553, 554, 564, 574, 577, 586, 595, 604,
	613, 622, 631, 634, 640, 650, 660, 670,
	672, 675, 678, 679, 688, 712, 736, 758,
	772, 786, 800, 814, 828, 843, 845, 849,
	853, 855, 870, 883, 897, 911, 925, 939,
	961, 962, 963, 966, 969, 970, 983, 985,
	988, 991, 994, 995, 997, 1004, 1006, 1014,
	1016, 1017, 1018, 1020, 1022, 1024, 1026, 1028,
	1030, 1032, 1034, 1036, 1038, 1041, 1044, 1046,
	1054, 1055, 1057, 1059, 1060, 1073, 1087, 1100,
	1114, 1135, 1149, 1171, 1173, 1174, 1196, 1218,
	1232, 1239, 1249, 1251, 1254, 1256, 1259, 1261,
	1264, 1267, 1270, 1271, 1274, 1275, 1285, 1295,
	1298, 1307, 1316, 1325, 1334, 1343, 1352, 1355,
	1361, 1371, 1381, 1391, 1393, 1396, 1399, 1400,
	1409, 1433, 1457, 1479, 1493, 1507, 1521, 1535,
	1549, 1564, 1566, 1570, 1574, 1576, 1591, 1604,
	1618, 1632, 1646, 1660, 1682, 1683, 1684, 1687,
	1690, 1691, 1696, 1697, 1700, 1703, 1704, 1708,
	1709, 1712, 1715, 1716, 1722, 1723, 1726, 1729,
	1750, 1771, 1792, 1807, 1814, 1823, 1825, 1828,
	1830, 1833, 1835, 1838, 1841, 1842, 1845, 1846,
	1849, 1850, 1859, 1868, 1870, 1878, 1886, 1894,
	1902, 1910, 1918, 1920, 1926, 1935, 1944, 1953,
	1955, 1958, 1961, 1962, 1970, 1971, 1977, 1979,
	1985, 1989, 1994,
}

var _sdp_trans_keys []byte = []byte{
	118, 61, 48, 57, 13, 10, 111, 61,
	127, 0, 32, 32, 127, 0, 31, 48,
	57, 32, 48, 57, 48, 57, 32, 48,
	57, 33, 35, 39, 42, 43, 45, 46,
	48, 57, 65, 90, 94, 126, 32, 33,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 90, 94, 126, 33, 35, 39, 42,
	43, 45, 46, 48, 57, 65, 90, 94,
	126, 32, 33, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 33,
	58, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 70, 71, 90, 94, 96, 97,
	102, 103, 126, 13, 33, 35, 39, 42,
	43, 45, 46, 48, 57, 65, 90, 94,
	126, 10, 115, 61, 0, 10, 13, 0,
	10, 13, 10, 98, 99, 101, 105, 112,
	116, 117, 61, 33, 35, 39, 42, 43,
	45, 46, 48, 57, 65, 90, 94, 126,
	33, 58, 35, 39, 42, 43, 45, 46,
	48, 57, 65, 90, 94, 126, 48, 57,
	13, 48, 57, 10, 98, 116, 61, 48,
	57, 32, 48, 57, 48, 57, 13, 48,
	57, 10, 61, 33, 35, 39, 42, 43,
	45, 46, 48, 57, 65, 90, 94, 126,
	13, 33, 58, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 10,
	61, 33, 35, 39, 42, 43, 45, 46,
	48, 57, 65, 90, 94, 126, 32, 33,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 90, 94, 126, 48, 57, 32, 47,
	48, 57, 33, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 32,
	33, 47, 35, 39, 42, 43, 45, 57,
	65, 90, 94, 126, 33, 35, 39, 42,
	43, 45, 46, 48, 57, 65, 90, 94,
	126, 13, 32, 33, 35, 39, 42, 43,
	45, 46, 48, 57, 65, 90, 94, 126,
	10, 61, 33, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 33,
	58, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 90, 94, 126, 48, 57, 13,
	48, 57, 10, 61, 0, 10, 13, 0,
	10, 13, 61, 33, 35, 39, 42, 43,
	45, 46, 48, 57, 65, 90, 94, 126,
	32, 33, 35, 39, 42, 43, 45, 46,
	48, 57, 65, 90, 94, 126, 33, 35,
	39, 42, 43, 45, 46, 48, 57, 65,
	90, 94, 126, 32, 33, 35, 39, 42,
	43, 45, 46, 48, 57, 65, 90, 94,
	126, 33, 50, 58, 35, 39, 42, 43,
	45, 46, 48, 57, 65, 70, 71, 90,
	94, 96, 97, 102, 103, 126, 13, 33,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 90, 94, 126, 13, 33, 47, 58,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 70, 71, 90, 94, 96, 97, 102,
	103, 126, 48, 57, 13, 13, 33, 47,
	58, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 70, 71, 90, 94, 96, 97,
	102, 103, 126, 13, 33, 47, 58, 35,
	39, 42, 43, 45, 46, 48, 57, 65,
	70, 71, 90, 94, 96, 97, 102, 103,
	126, 13, 33, 47, 58, 35, 39, 42,
	43, 45, 57, 65, 90, 94, 126, 58,
	48, 57, 65, 70, 97, 102, 13, 46,
	47, 58, 48, 57, 65, 70, 97, 102,
	48, 57, 46, 48, 57, 48, 57, 46,
	48, 57, 48, 57, 13, 48, 57, 13,
	48, 57, 46, 48, 57, 46, 46, 48,
	57, 46, 13, 46, 47, 58, 48, 57,
	65, 70, 97, 102, 13, 46, 47, 58,
	48, 57, 65, 70, 97, 102, 13, 47,
	58, 13, 47, 58, 48, 57, 65, 70,
	97, 102, 13, 47, 58, 48, 57, 65,
	70, 97, 102, 13, 47, 58, 48, 57,
	65, 70, 97, 102, 13, 47, 58, 48,
	57, 65, 70, 97, 102, 13, 47, 58,
	48, 57, 65, 70, 97, 102, 13, 47,
	58, 48, 57, 65, 70, 97, 102, 13,
	47, 58, 48, 57, 65, 70, 97, 102,
	13, 46, 47, 58, 48, 57, 65, 70,
	97, 102, 13, 46, 47, 58, 48, 57,
	65, 70, 97, 102, 13, 46, 47, 58,
	48, 57, 65, 70, 97, 102, 48, 57,
	46, 48, 57, 46, 48, 57, 46, 13,
	47, 58, 48, 57, 65, 70, 97, 102,
	13, 33, 47, 50, 51, 58, 35, 39,
	42, 43, 45, 46, 48, 57, 65, 70,
	71, 90, 94, 96, 97, 102, 103, 126,
	13, 33, 47, 58, 35, 39, 42, 43,
	45, 46, 48, 51, 52, 57, 65, 70,
	71, 90, 94, 96, 97, 102, 103, 126,
	13, 33, 45, 46, 47, 58, 35, 39,
	42, 43, 48, 57, 65, 70, 71, 90,
	94, 96, 97, 102, 103, 126, 13, 33,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 90, 94, 126, 13, 33, 45, 46,
	35, 39, 42, 43, 48, 57, 65, 90,
	94, 126, 13, 33, 35, 39, 42, 43,
	45, 46, 48, 57, 65, 90, 94, 126,
	13, 33, 45, 46, 35, 39, 42, 43,
	48, 57, 65, 90, 94, 126, 13, 33,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 90, 94, 126, 13, 33, 47, 35,
	39, 42, 43, 45, 46, 48, 57, 65,
	90, 94, 126, 48, 57, 13, 47, 48,
	57, 13, 47, 48, 57, 13, 47, 13,
	33, 47, 35, 39, 42, 43, 45, 46,
	48, 57, 65, 90, 94, 126, 13, 33,
	47, 35, 39, 42, 43, 45, 57, 65,
	90, 94, 126, 13, 33, 45, 46, 35,
	39, 42, 43, 48, 57, 65, 90, 94,
	126, 13, 33, 45, 46, 35, 39, 42,
	43, 48, 57, 65, 90, 94, 126, 13,
	33, 45, 46, 35, 39, 42, 43, 48,
	57, 65, 90, 94, 126, 13, 33, 45,
	46, 35, 39, 42, 43, 48, 57, 65,
	90, 94, 126, 13, 33, 47, 58, 35,
	39, 42, 43, 45, 46, 48, 57, 65,
	70, 71, 90, 94, 96, 97, 102, 103,
	126, 58, 61, 0, 10, 13, 0, 10,
	13, 10, 33, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 48,
	57, 32, 48, 57, 0, 10, 13, 0,
	10, 13, 61, 48, 57, 32, 100, 104,
	109, 115, 48, 57, 48, 57, 13, 32,
	100, 104, 109, 115, 48, 57, 13, 32,
	32, 61, 49, 57, 48, 57, 48, 57,
	48, 57, 48, 57, 48, 57, 48, 57,
	48, 57, 48, 57, 48, 57, 32, 48,
	57, 45, 48, 57, 48, 57, 13, 32,
	100, 104, 109, 115, 48, 57, 10, 49,
	57, 13, 32, 61, 33, 35, 39, 42,
	43, 45, 46, 48, 57, 65, 90, 94,
	126, 32, 33, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 33,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 90, 94, 126, 32, 33, 35, 39,
	42, 43, 45, 46, 48, 57, 65, 90,
	94, 126, 33, 50, 58, 35, 39, 42,
	43, 45, 46, 48, 57, 65, 70, 71,
	90, 94, 96, 97, 102, 103, 126, 13,
	33, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 90, 94, 126, 13, 33, 47,
	58, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 70, 71, 90, 94, 96, 97,
	102, 103, 126, 48, 57, 13, 13, 33,
	47, 58, 35, 39, 42, 43, 45, 46,
	48, 57, 65, 70, 71, 90, 94, 96,
	97, 102, 103, 126, 13, 33, 47, 58,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 70, 71, 90, 94, 96, 97, 102,
	103, 126, 13, 33, 47, 58, 35, 39,
	42, 43, 45, 57, 65, 90, 94, 126,
	58, 48, 57, 65, 70, 97, 102, 13,
	46, 47, 58, 48, 57, 65, 70, 97,
	102, 48, 57, 46, 48, 57, 48, 57,
	46, 48, 57, 48, 57, 13, 48, 57,
	13, 48, 57, 46, 48, 57, 46, 46,
	48, 57, 46, 13, 46, 47, 58, 48,
	57, 65, 70, 97, 102, 13, 46, 47,
	58, 48, 57, 65, 70, 97, 102, 13,
	47, 58, 13, 47, 58, 48, 57, 65,
	70, 97, 102, 13, 47, 58, 48, 57,
	65, 70, 97, 102, 13, 47, 58, 48,
	57, 65, 70, 97, 102, 13, 47, 58,
	48, 57, 65, 70, 97, 102, 13, 47,
	58, 48, 57, 65, 70, 97, 102, 13,
	47, 58, 48, 57, 65, 70, 97, 102,
	13, 47, 58, 48, 57, 65, 70, 97,
	102, 13, 46, 47, 58, 48, 57, 65,
	70, 97, 102, 13, 46, 47, 58, 48,
	57, 65, 70, 97, 102, 13, 46, 47,
	58, 48, 57, 65, 70, 97, 102, 48,
	57, 46, 48, 57, 46, 48, 57, 46,
	13, 47, 58, 48, 57, 65, 70, 97,
	102, 13, 33, 47, 50, 51, 58, 35,
	39, 42, 43, 45, 46, 48, 57, 65,
	70, 71, 90, 94, 96, 97, 102, 103,
	126, 13, 33, 47, 58, 35, 39, 42,
	43, 45, 46, 48, 51, 52, 57, 65,
	70, 71, 90, 94, 96, 97, 102, 103,
	126, 13, 33, 45, 46, 47, 58, 35,
	39, 42, 43, 48, 57, 65, 70, 71,
	90, 94, 96, 97, 102, 103, 126, 13,
	33, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 90, 94, 126, 13, 33, 45,
	46, 35, 39, 42, 43, 48, 57, 65,
	90, 94, 126, 13, 33, 35, 39, 42,
	43, 45, 46, 48, 57, 65, 90, 94,
	126, 13, 33, 45, 46, 35, 39, 42,
	43, 48, 57, 65, 90, 94, 126, 13,
	33, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 90, 94, 126, 13, 33, 47,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 90, 94, 126, 48, 57, 13, 47,
	48, 57, 13, 47, 48, 57, 13, 47,
	13, 33, 47, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 13,
	33, 47, 35, 39, 42, 43, 45, 57,
	65, 90, 94, 126, 13, 33, 45, 46,
	35, 39, 42, 43, 48, 57, 65, 90,
	94, 126, 13, 33, 45, 46, 35, 39,
	42, 43, 48, 57, 65, 90, 94, 126,
	13, 33, 45, 46, 35, 39, 42, 43,
	48, 57, 65, 90, 94, 126, 13, 33,
	45, 46, 35, 39, 42, 43, 48, 57,
	65, 90, 94, 126, 13, 33, 47, 58,
	35, 39, 42, 43, 45, 46, 48, 57,
	65, 70, 71, 90, 94, 96, 97, 102,
	103, 126, 58, 61, 0, 10, 13, 0,
	10, 13, 10, 98, 99, 101, 112, 116,
	61, 0, 10, 13, 0, 10, 13, 10,
	98, 99, 112, 116, 61, 0, 10, 13,
	0, 10, 13, 10, 98, 99, 101, 112,
	116, 117, 61, 0, 10, 13, 0, 10,
	13, 13, 33, 58, 35, 39, 42, 43,
	45, 46, 48, 57, 65, 70, 71, 90,
	94, 96, 97, 102, 103, 126, 13, 33,
	58, 35, 39, 42, 43, 45, 46, 48,
	57, 65, 70, 71, 90, 94, 96, 97,
	102, 103, 126, 13, 33, 58, 35, 39,
	42, 43, 45, 46, 48, 57, 65, 70,
	71, 90, 94, 96, 97, 102, 103, 126,
	13, 33, 58, 35, 39, 42, 43, 45,
	46, 48, 57, 65, 90, 94, 126, 58,
	48, 57, 65, 70, 97, 102, 13, 46,
	58, 48, 57, 65, 70, 97, 102, 48,
	57, 46, 48, 57, 48, 57, 46, 48,
	57, 48, 57, 13, 48, 57, 13, 48,
	57, 13, 46, 48, 57, 46, 46, 48,
	57, 46, 13, 46, 58, 48, 57, 65,
	70, 97, 102, 13, 46, 58, 48, 57,
	65, 70, 97, 102, 13, 58, 13, 58,
	48, 57, 65, 70, 97, 102, 13, 58,
	48, 57, 65, 70, 97, 102, 13, 58,
	48, 57, 65, 70, 97, 102, 13, 58,
	48, 57, 65, 70, 97, 102, 13, 58,
	48, 57, 65, 70, 97, 102, 13, 58,
	48, 57, 65, 70, 97, 102, 13, 58,
	48, 57, 65, 70, 97, 102, 13, 46,
	58, 48, 57, 65, 70, 97, 102, 13,
	46, 58, 48, 57, 65, 70, 97, 102,
	13, 46, 58, 48, 57, 65, 70, 97,
	102, 48, 57, 46, 48, 57, 46, 48,
	57, 46, 13, 58, 48, 57, 65, 70,
	97, 102, 58, 97, 107, 109, 114, 116,
	122, 97, 109, 97, 98, 99, 105, 107,
	109, 97, 98, 107, 109, 97, 98, 99,
	107, 109, 97, 107, 109,
}

var _sdp_single_lengths []byte = []byte{
	0, 1, 1, 0, 1, 1, 1, 1,
	1, 2, 0, 1, 0, 1, 1, 2,
	1, 2, 2, 2, 1, 1, 1, 3,
	3, 1, 7, 1, 1, 2, 0, 1,
	1, 2, 1, 0, 1, 0, 1, 1,
	1, 1, 3, 1, 1, 1, 2, 0,
	2, 1, 3, 1, 3, 1, 1, 1,
	2, 0, 1, 1, 1, 3, 3, 1,
	1, 2, 1, 2, 3, 2, 4, 0,
	1, 4, 4, 4, 1, 4, 0, 1,
	0, 1, 0, 1, 1, 1, 1, 1,
	1, 4, 4, 3, 3, 3, 3, 3,
	3, 3, 3, 0, 4, 4, 4, 0,
	1, 1, 1, 3, 6, 4, 6, 2,
	4, 2, 4, 2, 3, 0, 2, 2,
	2, 3, 3, 4, 4, 4, 4, 4,
	1, 1, 3, 3, 1, 1, 0, 1,
	3, 3, 1, 0, 5, 0, 6, 2,
	1, 1, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, 1, 0, 6,
	1, 0, 2, 1, 1, 2, 1, 2,
	3, 2, 4, 0, 1, 4, 4, 4,
	1, 4, 0, 1, 0, 1, 0, 1,
	1, 1, 1, 1, 1, 4, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 0,
	4, 4, 4, 0, 1, 1, 1, 3,
	6, 4, 6, 2, 4, 2, 4, 2,
	3, 0, 2, 2, 2, 3, 3, 4,
	4, 4, 4, 4, 1, 1, 3, 3,
	1, 5, 1, 3, 3, 1, 4, 1,
	3, 3, 1, 6, 1, 3, 3, 3,
	3, 3, 3, 1, 3, 0, 1, 0,
	1, 0, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 2, 2, 2, 2, 2,
	2, 2, 2, 0, 3, 3, 3, 0,
	1, 1, 1, 2, 1, 6, 2, 6,
	4, 5, 3,
}

var _sdp_range_lengths []byte = []byte{
	0, 0, 0, 1, 0, 0, 0, 0,
	1, 1, 1, 1, 1, 1, 6, 6,
	6, 6, 9, 6, 0, 0, 0, 0,
	0, 0, 0, 0, 6, 6, 1, 1,
	0, 0, 0, 1, 1, 1, 1, 0,
	0, 6, 6, 0, 0, 6, 6, 1,
	1, 6, 5, 6, 6, 0, 0, 6,
	6, 1, 1, 0, 0, 0, 0, 0,
	6, 6, 6, 6, 9, 6, 9, 1,
	0, 9, 9, 5, 3, 3, 1, 1,
	1, 1, 1, 1, 1, 1, 0, 1,
	0, 3, 3, 0, 3, 3, 3, 3,
	3, 3, 0, 3, 3, 3, 3, 1,
	1, 1, 0, 3, 9, 10, 8, 6,
	5, 6, 5, 6, 6, 1, 1, 1,
	0, 6, 5, 5, 5, 5, 5, 9,
	0, 0, 0, 0, 0, 6, 1, 1,
	0, 0, 0, 1, 1, 1, 1, 0,
	0, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1,
	0, 1, 0, 0, 6, 6, 6, 6,
	9, 6, 9, 1, 0, 9, 9, 5,
	3, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 1, 0, 3, 3, 0,
	3, 3, 3, 3, 3, 3, 0, 3,
	3, 3, 3, 1, 1, 1, 0, 3,
	9, 10, 8, 6, 5, 6, 5, 6,
	6, 1, 1, 1, 0, 6, 5, 5,
	5, 5, 5, 9, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9,
	9, 9, 6, 3, 3, 1, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1,
	0, 3, 3, 0, 3, 3, 3, 3,
	3, 3, 0, 3, 3, 3, 3, 1,
	1, 1, 0, 3, 0, 0, 0, 0,
	0, 0, 0,
}

var _sdp_index_offsets []int16 = []int16{
	0, 0, 2, 4, 6, 8, 10, 12,
	14, 17, 21, 23, 26, 28, 31, 39,
	48, 56, 65, 77, 86, 88, 90, 92,
	96, 100, 102, 110, 112, 120, 129, 131,
	134, 136, 139, 141, 143, 146, 148, 151,
	153, 155, 163, 173, 175, 177, 185, 194,
	196, 200, 208, 217, 225, 235, 237, 239,
	247, 256, 258, 261, 263, 265, 269, 273,
	275, 283, 292, 300, 309, 322, 331, 345,
	347, 349, 363, 377, 387, 392, 400, 402,
	405, 407, 410, 412, 415, 418, 421, 423,
	426, 428, 436, 444, 448, 455, 462, 469,
	476, 483, 490, 494, 498, 506, 514, 522,
	524, 527, 530, 532, 539, 555, 570, 585,
	594, 604, 613, 623, 632, 642, 644, 648,
	652, 655, 665, 674, 684, 694, 704, 714,
	728, 730, 732, 736, 740, 742, 750, 752,
	755, 759, 763, 765, 767, 774, 776, 784,
	787, 789, 791, 793, 795, 797, 799, 801,
	803, 805, 807, 809, 811, 814, 817, 819,
	827, 829, 831, 834, 836, 844, 853, 861,
	870, 883, 892, 906, 908, 910, 924, 938,
	948, 953, 961, 963, 966, 968, 971, 973,
	976, 979, 982, 984, 987, 989, 997, 1005,
	1009, 1016, 1023, 1030, 1037, 1044, 1051, 1055,
	1059, 1067, 1075, 1083, 1085, 1088, 1091, 1093,
	1100, 1116, 1131, 1146, 1155, 1165, 1174, 1184,
	1193, 1203, 1205, 1209, 1213, 1216, 1226, 1235,
	1245, 1255, 1265, 1275, 1289, 1291, 1293, 1297,
	1301, 1303, 1309, 1311, 1315, 1319, 1321, 1326,
	1328, 1332, 1336, 1338, 1345, 1347, 1351, 1355,
	1368, 1381, 1394, 1404, 1409, 1416, 1418, 1421,
	1423, 1426, 1428, 1431, 1434, 1436, 1439, 1441,
	1444, 1446, 1453, 1460, 1463, 1469, 1475, 1481,
	1487, 1493, 1499, 1502, 1506, 1513, 1520, 1527,
	1529, 1532, 1535, 1537, 1543, 1545, 1552, 1555,
	1562, 1567, 1573,
}

var _sdp_indicies []int16 = []int16{
	0, 1, 2, 1, 3, 1, 4, 1,
	5, 1, 6, 1, 7, 1, 1, 1,
	8, 9, 1, 1, 10, 11, 1, 12,
	13, 1, 14, 1, 15, 16, 1, 17,
	17, 17, 17, 17, 17, 17, 1, 18,
	19, 19, 19, 19, 19, 19, 19, 1,
	20, 20, 20, 20, 20, 20, 20, 1,
	21, 22, 22, 22, 22, 22, 22, 22,
	1, 23, 25, 23, 23, 23, 24, 24,
	23, 23, 24, 23, 1, 26, 27, 27,
	27, 27, 27, 27, 27, 1, 28, 1,
	29, 1, 30, 1, 1, 1, 1, 31,
	1, 1, 33, 32, 34, 1, 35, 36,
	37, 38, 39, 40, 41, 1, 42, 1,
	43, 43, 43, 43, 43, 43, 43, 1,
	44, 45, 44, 44, 44, 44, 44, 44,
	1, 46, 1, 47, 48, 1, 49, 1,
	35, 40, 1, 50, 1, 51, 1, 52,
	53, 1, 54, 1, 55, 56, 1, 57,
	1, 58, 1, 59, 59, 59, 59, 59,
	59, 59, 1, 60, 61, 62, 61, 61,
	61, 61, 61, 61, 1, 63, 1, 64,
	1, 65, 65, 65, 65, 65, 65, 65,
	1, 66, 67, 67, 67, 67, 67, 67,
	67, 1, 68, 1, 69, 70, 71, 1,
	72, 72, 72, 72, 72, 72, 72, 1,
	73, 74, 75, 74, 74, 74, 74, 74,
	1, 76, 76, 76, 76, 76, 76, 76,
	1, 77, 78, 76, 76, 76, 76, 76,
	76, 76, 1, 79, 1, 80, 1, 81,
	81, 81, 81, 81, 81, 81, 1, 82,
	83, 82, 82, 82, 82, 82, 82, 1,
	84, 1, 85, 86, 1, 87, 1, 88,
	1, 1, 1, 1, 89, 1, 1, 91,
	90, 92, 1, 93, 93, 93, 93, 93,
	93, 93, 1, 94, 95, 95, 95, 95,
	95, 95, 95, 1, 96, 96, 96, 96,
	96, 96, 96, 1, 97, 98, 98, 98,
	98, 98, 98, 98, 1, 99, 101, 102,
	99, 99, 99, 100, 100, 99, 99, 100,
	99, 1, 103, 104, 104, 104, 104, 104,
	104, 104, 1, 103, 104, 105, 107, 104,
	104, 104, 106, 106, 104, 104, 106, 104,
	1, 108, 1, 103, 1, 103, 104, 105,
	107, 104, 104, 104, 109, 109, 104, 104,
	109, 104, 1, 103, 104, 105, 107, 104,
	104, 104, 110, 110, 104, 104, 110, 104,
	1, 103, 104, 105, 107, 104, 104, 104,
	104, 104, 1, 112, 111, 113, 113, 1,
	103, 114, 105, 107, 115, 116, 116, 1,
	117, 1, 118, 119, 1, 120, 1, 121,
	122, 1, 123, 1, 103, 124, 1, 103,
	108, 1, 121, 125, 1, 121, 1, 118,
	126, 1, 118, 1, 103, 114, 105, 107,
	127, 128, 128, 1, 103, 114, 105, 107,
	129, 129, 129, 1, 103, 105, 107, 1,
	103, 105, 107, 129, 129, 129, 1, 103,
	105, 107, 128, 128, 128, 1, 103, 105,
	131, 130, 130, 130, 1, 103, 105, 133,
	132, 132, 132, 1, 103, 105, 133, 134,
	134, 134, 1, 103, 105, 133, 135, 135,
	135, 1, 103, 105, 133, 1, 136, 130,
	130, 1, 103, 114, 105, 133, 137, 132,
	132, 1, 103, 114, 105, 133, 138, 134,
	134, 1, 103, 114, 105, 133, 135, 135,
	135, 1, 139, 1, 114, 140, 1, 114,
	141, 1, 114, 1, 103, 105, 107, 116,
	116, 116, 1, 103, 104, 105, 142, 143,
	107, 104, 104, 104, 106, 106, 104, 104,
	106, 104, 1, 103, 104, 105, 107, 104,
	104, 104, 109, 144, 109, 104, 104, 109,
	104, 1, 103, 104, 104, 145, 105, 107,
	104, 104, 110, 110, 104, 104, 110, 104,
	1, 103, 104, 104, 104, 104, 146, 104,
	104, 1, 103, 104, 104, 147, 104, 104,
	148, 104, 104, 1, 103, 104, 104, 104,
	104, 149, 104, 104, 1, 103, 104, 104,
	150, 104, 104, 151, 104, 104, 1, 103,
	104, 104, 104, 104, 152, 104, 104, 1,
	103, 104, 153, 104, 104, 104, 154, 104,
	104, 1, 155, 1, 103, 105, 156, 1,
	103, 105, 157, 1, 103, 105, 1, 103,
	104, 153, 104, 104, 104, 158, 104, 104,
	1, 103, 104, 153, 104, 104, 104, 104,
	104, 1, 103, 104, 104, 150, 104, 104,
	159, 104, 104, 1, 103, 104, 104, 150,
	104, 104, 104, 104, 104, 1, 103, 104,
	104, 147, 104, 104, 160, 104, 104, 1,
	103, 104, 104, 147, 104, 104, 104, 104,
	104, 1, 103, 104, 105, 107, 104, 104,
	104, 144, 109, 104, 104, 109, 104, 1,
	112, 1, 161, 1, 1, 1, 1, 162,
	1, 1, 164, 163, 165, 1, 74, 74,
	74, 74, 74, 74, 74, 1, 166, 1,
	167, 168, 1, 1, 1, 1, 169, 1,
	1, 171, 170, 172, 1, 173, 1, 174,
	176, 176, 176, 176, 175, 1, 177, 1,
	178, 174, 179, 179, 179, 179, 177, 1,
	178, 174, 1, 174, 1, 180, 1, 181,
	1, 182, 1, 183, 1, 184, 1, 185,
	1, 186, 1, 187, 1, 188, 1, 189,
	1, 190, 1, 191, 190, 1, 192, 193,
	1, 193, 1, 194, 195, 196, 196, 196,
	196, 193, 1, 197, 1, 198, 1, 194,
	195, 1, 199, 1, 200, 200, 200, 200,
	200, 200, 200, 1, 201, 202, 202, 202,
	202, 202, 202, 202, 1, 203, 203, 203,
	203, 203, 203, 203, 1, 204, 205, 205,
	205, 205, 205, 205, 205, 1, 206, 208,
	209, 206, 206, 206, 207, 207, 206, 206,
	207, 206, 1, 210, 211, 211, 211, 211,
	211, 211, 211, 1, 210, 211, 212, 214,
	211, 211, 211, 213, 213, 211, 211, 213,
	211, 1, 215, 1, 210, 1, 210, 211,
	212, 214, 211, 211, 211, 216, 216, 211,
	211, 216, 211, 1, 210, 211, 212, 214,
	211, 211, 211, 217, 217, 211, 211, 217,
	211, 1, 210, 211, 212, 214, 211, 211,
	211, 211, 211, 1, 219, 218, 220, 220,
	1, 210, 221, 212, 214, 222, 223, 223,
	1, 224, 1, 225, 226, 1, 227, 1,
	228, 229, 1, 230, 1, 210, 231, 1,
	210, 215, 1, 228, 232, 1, 228, 1,
	225, 233, 1, 225, 1, 210, 221, 212,
	214, 234, 235, 235, 1, 210, 221, 212,
	214, 236, 236, 236, 1, 210, 212, 214,
	1, 210, 212, 214, 236, 236, 236, 1,
	210, 212, 214, 235, 235, 235, 1, 210,
	212, 238, 237, 237, 237, 1, 210, 212,
	240, 239, 239, 239, 1, 210, 212, 240,
	241, 241, 241, 1, 210, 212, 240, 242,
	242, 242, 1, 210, 212, 240, 1, 243,
	237, 237, 1, 210, 221, 212, 240, 244,
	239, 239, 1, 210, 221, 212, 240, 245,
	241, 241, 1, 210, 221, 212, 240, 242,
	242, 242, 1, 246, 1, 221, 247, 1,
	221, 248, 1, 221, 1, 210, 212, 214,
	223, 223, 223, 1, 210, 211, 212, 249,
	250, 214, 211, 211, 211, 213, 213, 211,
	211, 213, 211, 1, 210, 211, 212, 214,
	211, 211, 211, 216, 251, 216, 211, 211,
	216, 211, 1, 210, 211, 211, 252, 212,
	214, 211, 211, 217, 217, 211, 211, 217,
	211, 1, 210, 211, 211, 211, 211, 253,
	211, 211, 1, 210, 211, 211, 254, 211,
	211, 255, 211, 211, 1, 210, 211, 211,
	211, 211, 256, 211, 211, 1, 210, 211,
	211, 257, 211, 211, 258, 211, 211, 1,
	210, 211, 211, 211, 211, 259, 211, 211,
	1, 210, 211, 260, 211, 211, 211, 261,
	211, 211, 1, 262, 1, 210, 212, 263,
	1, 210, 212, 264, 1, 210, 212, 1,
	210, 211, 260, 211, 211, 211, 265, 211,
	211, 1, 210, 211, 260, 211, 211, 211,
	211, 211, 1, 210, 211, 211, 257, 211,
	211, 266, 211, 211, 1, 210, 211, 211,
	257, 211, 211, 211, 211, 211, 1, 210,
	211, 211, 254, 211, 211, 267, 211, 211,
	1, 210, 211, 211, 254, 211, 211, 211,
	211, 211, 1, 210, 211, 212, 214, 211,
	211, 211, 251, 216, 211, 211, 216, 211,
	1, 219, 1, 268, 1, 1, 1, 1,
	269, 1, 1, 271, 270, 272, 1, 35,
	36, 37, 39, 40, 1, 273, 1, 1,
	1, 1, 274, 1, 1, 276, 275, 277,
	1, 35, 36, 39, 40, 1, 278, 1,
	1, 1, 1, 279, 1, 1, 281, 280,
	282, 1, 35, 36, 37, 39, 40, 41,
	1, 283, 1, 1, 1, 1, 284, 1,
	1, 286, 285, 26, 27, 288, 27, 27,
	27, 287, 287, 27, 27, 287, 27, 1,
	26, 27, 288, 27, 27, 27, 289, 289,
	27, 27, 289, 27, 1, 26, 27, 288,
	27, 27, 27, 290, 290, 27, 27, 290,
	27, 1, 26, 27, 288, 27, 27, 27,
	27, 27, 27, 1, 292, 291, 293, 293,
	1, 26, 294, 288, 295, 296, 296, 1,
	297, 1, 298, 299, 1, 300, 1, 301,
	302, 1, 303, 1, 26, 304, 1, 26,
	305, 1, 26, 1, 301, 306, 1, 301,
	1, 298, 307, 1, 298, 1, 26, 294,
	288, 308, 309, 309, 1, 26, 294, 288,
	310, 310, 310, 1, 26, 288, 1, 26,
	288, 310, 310, 310, 1, 26, 288, 309,
	309, 309, 1, 26, 312, 311, 311, 311,
	1, 26, 314, 313, 313, 313, 1, 26,
	314, 315, 315, 315, 1, 26, 314, 316,
	316, 316, 1, 26, 314, 1, 317, 311,
	311, 1, 26, 294, 314, 318, 313, 313,
	1, 26, 294, 314, 319, 315, 315, 1,
	26, 294, 314, 316, 316, 316, 1, 320,
	1, 294, 321, 1, 294, 322, 1, 294,
	1, 26, 288, 296, 296, 296, 1, 292,
	1, 323, 324, 325, 326, 40, 327, 1,
	323, 325, 1, 323, 328, 329, 330, 324,
	325, 1, 323, 328, 324, 325, 1, 323,
	328, 329, 324, 325, 1, 323, 324, 325,
	1,
}

var _sdp_trans_targs []int16 = []int16{
	2, 0, 3, 4, 5, 6, 7, 8,
	9, 10, 9, 11, 12, 11, 13, 14,
	13, 15, 16, 15, 17, 18, 17, 19,
	247, 284, 20, 19, 21, 22, 23, 24,
	24, 25, 26, 27, 163, 229, 239, 234,
	34, 244, 28, 29, 29, 30, 31, 32,
	31, 33, 35, 36, 37, 36, 38, 39,
	38, 285, 41, 42, 43, 42, 136, 286,
	45, 46, 47, 46, 48, 49, 134, 48,
	50, 51, 50, 133, 52, 53, 51, 287,
	55, 56, 56, 57, 58, 59, 58, 288,
	61, 62, 62, 43, 64, 65, 66, 65,
	67, 68, 67, 69, 70, 108, 128, 59,
	69, 71, 73, 76, 72, 74, 75, 77,
	94, 107, 78, 89, 93, 79, 80, 87,
	81, 82, 85, 83, 84, 86, 88, 90,
	92, 91, 95, 103, 96, 99, 97, 98,
	100, 101, 102, 104, 105, 106, 109, 127,
	110, 111, 112, 113, 125, 114, 115, 123,
	116, 117, 121, 118, 119, 120, 122, 124,
	126, 130, 131, 131, 132, 289, 135, 49,
	135, 137, 137, 43, 139, 140, 141, 140,
	144, 142, 39, 143, 146, 147, 148, 149,
	150, 151, 152, 153, 154, 155, 156, 157,
	158, 159, 160, 161, 162, 290, 147, 164,
	165, 166, 165, 167, 168, 167, 169, 170,
	208, 228, 32, 169, 171, 173, 176, 172,
	174, 175, 177, 194, 207, 178, 189, 193,
	179, 180, 187, 181, 182, 185, 183, 184,
	186, 188, 190, 192, 191, 195, 203, 196,
	199, 197, 198, 200, 201, 202, 204, 205,
	206, 209, 227, 210, 211, 212, 213, 225,
	214, 215, 223, 216, 217, 221, 218, 219,
	220, 222, 224, 226, 230, 231, 231, 232,
	233, 235, 236, 236, 237, 238, 240, 241,
	241, 242, 243, 245, 246, 246, 232, 248,
	251, 249, 250, 252, 270, 283, 253, 265,
	269, 254, 255, 263, 256, 257, 261, 258,
	259, 260, 262, 264, 266, 268, 267, 271,
	279, 272, 275, 273, 274, 276, 277, 278,
	280, 281, 282, 40, 60, 44, 138, 145,
	54, 63, 129,
}

var _sdp_trans_actions []byte = []byte{
	0, 0, 0, 1, 41, 0, 0, 0,
	1, 31, 0, 1, 33, 0, 1, 35,
	0, 1, 37, 0, 1, 39, 0, 1,
	1, 1, 43, 0, 0, 0, 0, 1,
	0, 45, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 1, 0, 25, 1, 27,
	0, 0, 0, 1, 19, 0, 1, 21,
	0, 0, 0, 1, 17, 0, 13, 0,
	0, 1, 55, 0, 1, 57, 57, 0,
	1, 63, 0, 0, 0, 61, 0, 0,
	0, 1, 0, 25, 1, 27, 0, 0,
	0, 1, 0, 29, 0, 1, 5, 0,
	1, 7, 0, 1, 1, 1, 1, 9,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 0, 3, 0, 1, 59,
	0, 1, 0, 15, 0, 1, 0, 0,
	0, 0, 23, 0, 0, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 53, 0, 0, 0, 0, 0,
	1, 5, 0, 1, 7, 0, 1, 1,
	1, 1, 9, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 1, 0, 49,
	0, 0, 1, 0, 51, 0, 0, 1,
	0, 3, 0, 0, 1, 0, 47, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 11, 0, 0,
	0, 0, 0,
}

const sdp_start int = 1
const sdp_first_final int = 285
const sdp_error int = 0

const sdp_en_main int = 1

//line parser.rl:8

// Parse scan and parse bytes array to SDP Message structure
func Parse(data []byte) (*Message, error) {
	var cs, p, pe, m int
	pe = len(data)

	// when media index is >= then it is media fields context
	// otherwise it is session context
	mediaIdx := -1
	msg := &Message{}

//line parser.rl:238

//line parser.go:752
	{
		cs = sdp_start
	}

//line parser.rl:240

//line parser.go:759
	{
		var _klen int
		var _trans int
		var _acts int
		var _nacts uint
		var _keys int
		if p == pe {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		_keys = int(_sdp_key_offsets[cs])
		_trans = int(_sdp_index_offsets[cs])

		_klen = int(_sdp_single_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + _klen - 1)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + ((_upper - _lower) >> 1)
				switch {
				case data[p] < _sdp_trans_keys[_mid]:
					_upper = _mid - 1
				case data[p] > _sdp_trans_keys[_mid]:
					_lower = _mid + 1
				default:
					_trans += int(_mid - int(_keys))
					goto _match
				}
			}
			_keys += _klen
			_trans += _klen
		}

		_klen = int(_sdp_range_lengths[cs])
		if _klen > 0 {
			_lower := int(_keys)
			var _mid int
			_upper := int(_keys + (_klen << 1) - 2)
			for {
				if _upper < _lower {
					break
				}

				_mid = _lower + (((_upper - _lower) >> 1) & ^1)
				switch {
				case data[p] < _sdp_trans_keys[_mid]:
					_upper = _mid - 2
				case data[p] > _sdp_trans_keys[_mid+1]:
					_lower = _mid + 2
				default:
					_trans += int((_mid - int(_keys)) >> 1)
					goto _match
				}
			}
			_trans += _klen
		}

	_match:
		_trans = int(_sdp_indicies[_trans])
		cs = int(_sdp_trans_targs[_trans])

		if _sdp_trans_actions[_trans] == 0 {
			goto _again
		}

		_acts = int(_sdp_trans_actions[_trans])
		_nacts = uint(_sdp_actions[_acts])
		_acts++
		for ; _nacts > 0; _nacts-- {
			_acts++
			switch _sdp_actions[_acts-1] {
			case 0:
//line parser.rl:21
				m = p
			case 1:
//line parser.rl:22

				if mediaIdx == -1 {
					msg.info = data[m:p]
				} else {
					i := len(msg.Medias) - 1
					msg.Medias[i].info = data[m:p]
				}

			case 2:
//line parser.rl:30

				if mediaIdx == -1 {
					msg.Conn.netType = data[m:p]
				} else {
					i := len(msg.Medias) - 1
					msg.Medias[i].Conn.netType = data[m:p]
				}

			case 3:
//line parser.rl:38

				if mediaIdx == -1 {
					msg.Conn.addrType = data[m:p]
				} else {
					i := len(msg.Medias) - 1
					msg.Medias[i].Conn.addrType = data[m:p]
				}

			case 4:
//line parser.rl:46

				if mediaIdx == -1 {
					msg.Conn.address = data[m:p]
				} else {
					i := len(msg.Medias) - 1
					msg.Medias[i].Conn.address = data[m:p]
				}

			case 5:
//line parser.rl:54

				if mediaIdx == -1 {
					msg.Medias = make(Medias, 1)
					mediaIdx = 0
				} else {
					mediaIdx++
					msg.Medias = append(msg.Medias, Media{})
				}

			case 6:
//line parser.rl:63

				if mediaIdx == -1 {
					msg.Attr = append(msg.Attr, Attribute{})
					i := len(msg.Attr) - 1
					msg.Attr[i].key = data[m:p]
				} else {
					msg.Medias[mediaIdx].Attr = append(msg.Medias[mediaIdx].Attr, Attribute{})
					i := len(msg.Medias[mediaIdx].Attr) - 1
					msg.Medias[mediaIdx].Attr[i].key = data[m:p]
				}

			case 7:
//line parser.rl:74

				if mediaIdx == -1 {
					i := len(msg.Attr) - 1
					msg.Attr[i].value = data[m:p]
				} else {
					i := len(msg.Medias[mediaIdx].Attr) - 1
					msg.Medias[mediaIdx].Attr[i].value = data[m:p]
				}

			case 8:
//line parser.rl:83

				if mediaIdx == -1 {
					msg.Attr = append(msg.Attr, Attribute{})
					i := len(msg.Attr) - 1
					msg.Attr[i].flag = data[m:p]
					msg.Attr[i].isFlag = true
				} else {
					msg.Medias[mediaIdx].Attr = append(msg.Medias[mediaIdx].Attr, Attribute{})
					i := len(msg.Medias[mediaIdx].Attr) - 1
					msg.Medias[mediaIdx].Attr[i].flag = data[m:p]
					msg.Medias[mediaIdx].Attr[i].isFlag = true
				}

			case 9:
//line parser.rl:96

				msg.Time = append(msg.Time, TimeDesc{start: data[m:p]})

			case 10:
//line parser.rl:99

				i := len(msg.Time) - 1
				msg.Time[i].stop = data[m:p]

			case 11:
//line parser.rl:103

				i := len(msg.Time) - 1
				msg.Time[i].Repeat = append(msg.Time[i].Repeat, data[m:p])

			case 12:
//line parser.rl:107

				if mediaIdx == -1 {
					msg.BandWidth = append(msg.BandWidth, BandWidth{bt: data[m:p]})
				} else {
					i := len(msg.Medias) - 1
					msg.Medias[i].BandWidth = append(msg.Medias[i].BandWidth, BandWidth{bt: data[m:p]})
				}

			case 13:
//line parser.rl:115

				if mediaIdx == -1 {
					i := len(msg.BandWidth) - 1
					msg.BandWidth[i].bw = data[m:p]
				} else {
					i := len(msg.Medias) - 1
					j := len(msg.Medias[i].BandWidth) - 1
					msg.Medias[i].BandWidth[j].bw = data[m:p]
				}

			case 14:
//line parser.rl:125

				if mediaIdx == -1 {
					msg.encKey = data[m:p]
				} else {
					i := len(msg.Medias) - 1
					msg.Medias[i].encKey = data[m:p]
				}

			case 15:
//line parser.rl:158
				msg.Origin.username = data[m:p]
			case 16:
//line parser.rl:159
				msg.Origin.sessID = data[m:p]
			case 17:
//line parser.rl:160
				msg.Origin.sessVer = data[m:p]
			case 18:
//line parser.rl:161
				msg.Origin.netType = data[m:p]
			case 19:
//line parser.rl:162
				msg.Origin.addrType = data[m:p]
			case 20:
//line parser.rl:181
				msg.ver = data[m]
			case 21:
//line parser.rl:184
				msg.Origin.unicAddr = data[m:p]
			case 22:
//line parser.rl:185
				msg.subject = data[m:p]
			case 23:
//line parser.rl:188
				msg.uri = data[m:p]
			case 24:
//line parser.rl:190
				msg.Email = append(msg.Email, data[m:p])
			case 25:
//line parser.rl:192
				msg.Phone = append(msg.Phone, data[m:p])
			case 26:
//line parser.rl:202
				msg.tzones = data[m:p]
			case 27:
//line parser.rl:207
				msg.Medias[mediaIdx].mtype = data[m:p]
			case 28:
//line parser.rl:208
				msg.Medias[mediaIdx].port = data[m:p]
			case 29:
//line parser.rl:209
				msg.Medias[mediaIdx].nport = data[m:p]
			case 30:
//line parser.rl:210
				msg.Medias[mediaIdx].proto = data[m:p]
			case 31:
//line parser.rl:211
				msg.Medias[mediaIdx].fmt = data[m:p]
//line parser.go:1031
			}
		}

	_again:
		if cs == 0 {
			goto _out
		}
		p++
		if p != pe {
			goto _resume
		}
	_test_eof:
		{
		}
	_out:
		{
		}
	}

//line parser.rl:241
	if cs >= sdp_first_final {
		return msg, nil
	}
	// improve error message
	e := pe
	if (p + 12) < e {
		e = p + 12
	}
	s := 0
	if (p - 24) > 0 {
		s = p - 24
	}
	return nil, ErrorSDPParsing.msg("%q... [position=%d]", data[s:e], p)
}
