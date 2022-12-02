<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Helloctlgateway;

/**
 */
class HelloCtlClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Helloctlgateway\SayHelloReq $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SayHello(\Helloctlgateway\SayHelloReq $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/helloctlgateway.HelloCtl/SayHello',
        $argument,
        ['\Helloctlgateway\SayHelloResp', 'decode'],
        $metadata, $options);
    }

}
