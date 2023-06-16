package org.example;

import audio.AudioServiceGrpc;
import audio.AudioServiceOuterClass;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class ClientListen {
    public static void main(String[] args) {

        ManagedChannel channel = ManagedChannelBuilder.forTarget("localhost:8012").usePlaintext().build();
        AudioServiceGrpc.AudioServiceBlockingStub stub = AudioServiceGrpc.newBlockingStub(channel);
        AudioServiceOuterClass.ListenRequest request=AudioServiceOuterClass.ListenRequest.newBuilder().setFileId("Criste").build();
        AudioServiceOuterClass.ListenResponse response= stub.listen(request);
        System.out.println(response);
        channel.shutdownNow();
//        ManagedChannel channel =  ManagedChannelBuilder.forTarget("localhost:8012").usePlaintext().build();
//
//        AudioServiceGrpc.AudioServiceBlockingStub stub= AudioServiceGrpc.newBlockingStub(channel);
//
//        AudioServiceOuterClass.ListenRequest request = AudioServiceOuterClass.ListenRequest.newBuilder().setFileId("Criste").build();
//
//
//        AudioServiceOuterClass.ListenResponse response = stub.listen(request);
//
//        System.out.println(response);
//        channel.shutdownNow();
    }
}
