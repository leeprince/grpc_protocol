<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/helloctlgateway/service_helloctlgateway.proto

namespace GPBMetadata\Protos\Helloctlgateway;

class ServiceHelloctlgateway
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Api\Annotations::initOnce();
        \GPBMetadata\Protos\Helloctlgateway\Hello::initOnce();
        $pool->internalAddGeneratedFile(
            '

4protos/helloctlgateway/service_helloctlgateway.protohelloctlgateway"protos/helloctlgateway/hello.proto2y
HelloCtlm
SayHello.helloctlgateway.SayHelloReq.helloctlgateway.SayHelloResp"$سل"/helloctlgateway/sayhello:*BZ./;helloctlgatewaybproto3'
        , true);

        static::$is_initialized = true;
    }
}

