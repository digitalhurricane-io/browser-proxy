##Browser Proxy

Great for making requests from the browser to servers without CORS support.

Also great if you need to make a request in the browser to an endpoint that doesn't support ssl.

###To Use:

A.) Run the app on a server. Personally I use google cloud run for this. 
You can use the deploy.cloudrun.sh script to deploy to cloud run.
Just replace occurrences of "my-project" in the script with the name of your project.

B.) Make a request to the proxy url and add a "target" header with the value set to the destination url.

The request will be proxied to the url in the "target" header.

So for example, you can add a "target" header with a value of "http://example.com/status?arg1=a" and the 
request will be proxied to example.com

