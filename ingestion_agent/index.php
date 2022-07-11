<?php
    $redis = new Redis();

    try {
        $redis->connect('127.0.0.1', 6379);
    } catch (RedisException $exception) {
        print_r($exception);
    }

    echo "Connection to server successful. ";
    echo "Server is running: ".$redis->ping();
    
?>