<?php
    $redis = new Redis();

    $redis->connect('127.0.0.1', 6379);

    echo "Connection to server successful\n";
    echo "Server is running: ".$redis->ping()."\n";

    // reads and stores incoming http POST request
    $json = file_get_contents("php://input");
    $json_arr = json_decode($json, true);

    if ($json_arr['endpoint'] && ['data']) {
        $postback = json_encode($json_arr);
        $push = $redis->rPush('data', $postback);
        if ($push) {
            echo "Pushed successfully\n";
        } else {
            print_r("Error pushing to Redis\n");
        }
    }
    else {
        echo "No request found\n";
    }
?>