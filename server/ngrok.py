from pyngrok import ngrok
import os

def clear():
    os.system("clear")

tun = ngrok.connect(8080, "tcp")
ngrok_process = ngrok.get_ngrok_process()

clear()

print("game url: ", tun.public_url[6:])

try:
    ngrok_process.proc.wait()
except KeyboardInterrupt:
    print(" Shutting down server.")
    ngrok.kill()

