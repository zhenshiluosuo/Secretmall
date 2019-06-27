<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no,viewport-fit=cover">
    <link rel="stylesheet" href="../static/css/search_style.css">
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
                    <td><img src="../static/img/item/{{.Descirption}}.jpg" width="150px" height="150px"/></td>
                    <td><table>
                        <tr><td>{{.Name}}<br/><br/></td></tr>
                        <tr><td style="color: red;">￥{{.Price}}</td></tr>
                        <tr><td style="font-size: 10px;color: darkgrey;"><u>{{.Name}}专卖店</u></td></tr>
                    </table></td>
                </tr>
                {{end}}
            </table>
        </div>
    </div>
</body>
</html>
