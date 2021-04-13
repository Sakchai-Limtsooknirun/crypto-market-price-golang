STEP TO DEPLOY TO HEROKU

-- Prepare Heroku

<code> heroku login </code>

-- Prepare Project

<p>Add below code for dynamic port </p>

<code>
func getPort() string {
     var port = os.Getenv("PORT")
     if port == "" {
        port = "8080"
        fmt.Println("No Port In Heroku" + port)
     }
     return ":" + port // ----> (B)
}
</code>

<p>ติดตั้ง(Install) GoDep</p>
<code>go get github.com/tools/godep</code>

<code>godeb save</code>

<p>สร้าง Procfile</p>

<code>touch Procfile</code>

<code>echo web: {{ project directory}} > Procfile</code>

<p> set default language </p>

<code> heroku buildpacks:set heroku/go  </code>

<p>run go mod</p>

<code> go mod init </code>

<code> go mod tidy </code>

<p> push all to github </p>

<p> Let deploy to heroku </p>

<code> heroku create </code>

<code> git push heroku main </code>

<code> heroku open </code>