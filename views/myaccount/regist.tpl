<!DOCTYPE html>

<html>
<head>
    <title>Beego</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no,viewport-fit=cover">
    <link rel="stylesheet" href="../static/css/login_style.css">
</head>
<body>
    <form action="/myaccount" method="post">
        <div style="margin-top: 80px;">
            <div class="title" style="margin-bottom: 50px;">
                <span><img src="../static/img/logo.jpg" align="center" height="140px" width="140px"></img></span>
            </div>
            <div class="main-agileinfo">
                <h2>注册用户</h2>
                    <input type="text" name="username" class="name" placeholder="请输入您的手机号" required="">
                    <input type="password" name="password" class="password" placeholder="请输入密码" required="">
                    <input type="password" name="password" class="password" placeholder="请确认密码" required="">
                    <ul>
                        <li>
                            <input type="checkbox" id="brand1" value="">
                            <label for="brand1"><span></span>记得我</label>
                        </li>
                    </ul>
                <a href="#" style="font-size: 16px; color:#F00"><b>查看本站用户协议</b>
                    </a><br/>
                    <div class="clear"></div>
                    <input type="submit" value="Submit">
            </div>
        </div>
    </form>
</body>
</html>
