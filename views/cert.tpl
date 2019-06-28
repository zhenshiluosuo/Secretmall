<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no,viewport-fit=cover">
    <link rel="stylesheet" href="../static/css/search_style.css">
</head>
<body>
    <div>
        <div>
          <div>
             <header>
               <h1>购物车</h1>
              </header>
          </div>
        </div>
    </div>
    <div>
        <div>
		     <table style="border: 2px solid #ff9000; border-radius:15px 15px 15px 15px; width: 100%">
                {{range .item}}
                <tr style="border-bottom:2px solid #ff9000;">
                    <td><img src="../static/img/item/{{.Descirption}}.jpg" width="150px" height="150px"/></td>
                    <td><table>
                        <tr><td>{{.Name}}<br/><br/></td></tr>
                        <tr><td style="color: red;">￥{{.Price}}</td></tr>
						<tr><td style="color: orange;">{{.Amount}}</td></tr>
                        <tr><td style="font-size: 10px;color: darkgrey;"><u>{{.Name}}专卖店</u></td></tr>
                    </table></td>
                </tr>
                {{end}}
            </table>
        </div>
    </div>
</body>
</html>

