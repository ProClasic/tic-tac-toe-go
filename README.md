TIC-TAC-TOE-GO

A terminal based online game made with go.

Installation:
  
1. You need to have python and golang installed on your device.
2. In your terminal go to the root directory of the game file.
3. Run "bash setup.sh" and let the installation process complete
4. After that a executable file will be generated in the root directory as "ttt" in linux/mac or "ttt.exe" on windows.

How to play:

*** Video explaination: https://www.youtube.com/watch?v=wkOkDPd9o0c ***

1. Run the executable file by "./ttt" or "./ttt.exe"
2. Enter your name

*** If you are Host ***

1. Select the option "Host new game" and hit enter in the main menu.
2. Then open an another terminal and go to the server directory.
3. In the server directory run "python server.py"

*Note*: If you see any errors, try checking your internet connection.
        Or if you are using a mobile device turn on: 
        - mobile hotspot if you are using mobile data
        - wifi tethering if you are using wifi

4. After that you will get a tcp link copy that and send it to your opponent and wait for the game.

*** If you want to join a game ***

1. You need to have a tcp game link to proceed further.
2. Just paste the tcp link in the "Join game" option.
