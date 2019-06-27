<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no,viewport-fit=cover">
    <link rel="stylesheet" href="../static/css/search_style.css">
    <style>
        .nav {
            width:100%;
            height:75px;
            margin:0 auto;
            background:#ffffff;
            position:fixed;
            bottom:0;
            text-align:center;
        }
    </style>
</head>
<body>
    <header>
        <form>
            <div id="container">
                <div class="search bar6">
                    <form>
                        <input type="text" placeholder="试试搜搜：在世真龙" name="cname">
                        <button type="submit"></button>
                    </form>
                </div>
            </div>
        </form>
    </header>
    <div>
        <div>
            <table style="border: 2px solid #ff9000; border-radius:15px 15px 15px 15px; width: 100%">
                {{range .item}}
                <tr style="border-bottom:2px solid #ff9000;">
<!--                    <td><img src="../static/img/item/{{.Descirption}}.jpg" width="150px" height="150px"/></td>-->
<!--                    <td><table>-->
<!--                        <tr><td>{{.Name}}<br/><br/></td></tr>-->
<!--                        <tr><td style="color: red;">￥{{.Price}}</td></tr>-->
<!--                        <tr><td style="font-size: 10px;color: darkgrey;"><u>{{.Name}}专卖店</u></td></tr>-->
<!--                    </table></td>-->
                    <td>
                    <table border="0">
                        <tr>
                            <td><img src="../static/img/item/{{.Descirption}}.jpg" width="150px" height="150px"/></td>
                            <td valign="top">
                                <table border="0">
                                <tr>
<!--                                    <td width="250"><font color="blue">HTML CSS JavaScript大全 </font></td>-->
                                    <td width="100"><font color="#dc143c" style="font-size: 12px; color: dimgrey;">{{.Name}}</font></td>
                                </tr>
                                <tr valign="top">
                                    <td ><font color="#dc143c" size="3">现价： {{.Price}}￥</font> </td>
                                </tr>
                                <tr>
                                    <td ><font style="color: #ff9000; font-size: 10px">七天无理由退换</font><font style="color: darkgrey">&nbsp;&nbsp;超市直营</font></td>
                                </tr>
                                <tr>
                                    <td><font color="blue">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{.Amount}}评论</font></td>
                                </tr>
                                <tr>
                                    <td><img src="../static/img/baoyou.png"/></td>
<!--                                    <td><img src="hot1.png"></td>-->
                                </tr>
                                </table>
                                <table border="0">
                                    <tr>
                                        <td bgcolor="#dc143c" align="center" width="80" style="font-size: 12px"><a href="#2" >加入购物车</a></td>
                                        <td width="25"></td>
                                        <td bgcolor="#ffb6c1" width="80" align="center"><font color="#dc143c">收藏</font></td>
                                    </tr>
                                </table>
                            </td>
                        </tr>
                    </table>
                    </td>
                </tr>
                {{end}}
            </table>
        </div>
    </div>


    <div class="nav">
        <div style="float: left; width: 24.5%; border-right: 1px rgb(255, 123, 0) solid;  "><a href="/"><img src="../static/img/x1.png"/></a></div>
        <div style="float: left; width: 24.5%; border-right: 1px rgb(255, 123, 0) solid;  "><a href="/cert"><img src="../static/img/x2.png"/></a></div>
        <div style="float: left; width: 24.5%; border-right: 1px rgb(255, 123, 0) solid;  "><img src="../static/img/x3.png"/></div>
        <div style="float: left; width: 24.5%;"><img src="../static/img/x4.png"/></div>
    </div>

</body>
</html>
