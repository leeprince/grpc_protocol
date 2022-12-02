<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Helloctlgateway;

/**
 */
class HelloCtlStub {

    /**
     * @param \Helloctlgateway\SayHelloReq $request client request
     * @param \Grpc\ServerContext $context server request context
     * @return \Helloctlgateway\SayHelloResp for response data, null if if error occured
     *     initial metadata (if any) and status (if not ok) should be set to $context
     */
    public function SayHello(
        \Helloctlgateway\SayHelloReq $request,
        \Grpc\ServerContext $context
    ): ?\Helloctlgateway\SayHelloResp {
        $context->setStatus(\Grpc\Status::unimplemented());
        return null;
    }

    /**
     * Get the method descriptors of the service for server registration
     *
     * @return array of \Grpc\MethodDescriptor for the service methods
     */
    public final function getMethodDescriptors(): array
    {
        return [
            '/helloctlgateway.HelloCtl/SayHello' => new \Grpc\MethodDescriptor(
                $this,
                'SayHello',
                '\Helloctlgateway\SayHelloReq',
                \Grpc\MethodDescriptor::UNARY_CALL
            ),
        ];
    }

}
