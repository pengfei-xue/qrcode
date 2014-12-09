<html>
<head>
<title>QR Generator</title>
<style type="text/css">
.container {
    margin: auto;
    width: 50%;
    padding-top: 100px;
}

.container img {
    margin-right: auto;
    margin-left: auto;
    margin-top: 30px;
    border: solid 1px;
    width: 50%;
    display: block;
}

.container form {
    margin: auto;
    width: 50%;
}

.container input[type=submit] {
    float: right;
}
</style>
</head>
<body>
    <div class="container">
        <form action="/qr" method="post">
            <input type="text" name="text"></input>
            <input type="submit" value="Get QR"></input>
        </form>
        <img src="data:image/png;base64,{{ .Img }}" alt="二维码"/>
    </div>
</body>
