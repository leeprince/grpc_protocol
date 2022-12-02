<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Helloctl;

/**
 */
class HelloCtlStub {

    /**
     * 说Hello
     * @param \Helloctl\SayHelloReq $request client request
     * @param \Grpc\ServerContext $context server request context
     * @return \Helloctl\SayHelloResp for response data, null if if error occured
     *     initial metadata (if any) and status (if not ok) should be set to $context
     */
    public function SayHello(
        \Helloctl\SayHelloReq $request,
        \Grpc\ServerContext $context
    ): ?\Helloctl\SayHelloResp {
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
            '/helloctl.HelloCtl/SayHello' => new \Grpc\MethodDescriptor(
                $this,
                'SayHello',
                '\Helloctl\SayHelloReq',
                \Grpc\MethodDescriptor::UNARY_CALL
            ),
        ];
    }

}
