# zmq-serialize

## Setup

#### System update
    
    sudo apt-get update && sudo apt-get install git
        
#### Install libzmq 3.2.5

    sudo apt-get install build-essential libtool autoconf automake uuid-dev
    wget http://download.zeromq.org/zeromq-3.2.5.tar.gz
    tar xvzf zeromq-3.2.5.tar.gz
    cd zeromq-3.2.5
    ./configure
    make
    sudo make install
    sudo ldconfig
    #sudo ln -sf /usr/local/lib/libzmq.so /usr/lib/libzmq.so
  
##### Install node.js 4.x & zmq binding

    curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash -
    sudo apt-get install -y nodejs
    sudo npm install -g zmq

#### Install golang 1.6 & zmq binding

    sudo apt-get install pkg-config
    curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
    tar -xvf go1.6.linux-amd64.tar.gz
    sudo mv go /usr/local
    nano ~/.profile
    export PATH=$PATH:/usr/local/go/bin