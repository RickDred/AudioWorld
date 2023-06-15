package org.example;
import com.google.protobuf.ByteString;
import audio.AudioServiceGrpc;
import audio.AudioServiceOuterClass;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import java.nio.file.Files;
import java.nio.file.Path;

public class Client {
    public static void main(String[] args) {
        try {
            // Путь к аудиофайлу
            Path filePath = Path.of("src/main/newAudio/Кристина Ашмарина - Blue Bird (Naruto Shippuden Opening) (dizer.net).mp3");

            // Чтение содержимого файла в виде байтов
            byte[] audioBytes = Files.readAllBytes(filePath);
            ByteString audioData = ByteString.copyFrom(audioBytes);

            // Дальнейшая обработка байтов аудио (например, передача по сети, сохранение в базе данных и т. д.)
            ManagedChannel channel =  ManagedChannelBuilder.forTarget("localhost:8011").usePlaintext().build();

            AudioServiceGrpc.AudioServiceBlockingStub stub= AudioServiceGrpc.newBlockingStub(channel);

            AudioServiceOuterClass.UploadRequest request = AudioServiceOuterClass.UploadRequest.newBuilder().setAudioData(audioData).setName("Criste").build();
            // ...

            AudioServiceOuterClass.UploadResponse response = stub.upload(request);

            System.out.println(response);
            channel.shutdownNow();
        } catch (Exception e) {
            e.printStackTrace();
        }



    }
}
