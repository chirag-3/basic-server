This is a server

#The hello path working
When someone hits /hello
the helloHandler gets executed, which result i sthe applicaetion sending back a response which has plain string.
The response also has headers like COntent-type that is read by the browser to detect that respone is just plain text, And hence it just deisplay it. The COntent-type For that is returned as text
