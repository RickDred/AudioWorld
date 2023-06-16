package org.example;
import audio.AudioServiceGrpc;
import audio.AudioServiceOuterClass;
import com.google.protobuf.ByteString;
import com.mongodb.MongoClient;
import com.mongodb.MongoClientURI;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import org.bson.Document;

import javax.sound.sampled.*;
import javax.sound.sampled.AudioInputStream;
import javax.sound.sampled.AudioSystem;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.URL;

public class ControllerServiceImpl extends AudioServiceGrpc.AudioServiceImplBase {
    // Создание подключения к MongoDB
//    MongoClientURI connectionString = new MongoClientURI("mongodb://atlas-sql-648b67ce6d0c983b4099ed12-lohh9.a.query.mongodb.net/name?ssl=true&authSource=admin");
//    MongoClient mongoClient = new MongoClient(connectionString);
//
//    // Получение базы данных
//    MongoDatabase database = mongoClient.getDatabase("name");


//
@Override
public void listen(AudioServiceOuterClass.ListenRequest request, StreamObserver<AudioServiceOuterClass.ListenResponse> responseStreamObserver) {
    String audioUrl = request.getFileId();

    try {
        // Загрузка аудиофайла по URL
        InputStream inputStream = ControllerServiceImpl.class.getResourceAsStream("src/main/java/org/example/Book/" + audioUrl + ".mp3");
        if (inputStream == null) {
            throw new FileNotFoundException("Audio file not found: " + audioUrl);
        }

        AudioInputStream audioInputStream = AudioSystem.getAudioInputStream(inputStream);

        // Получение Clip из AudioInputStream
        Clip clip = AudioSystem.getClip();
        clip.open(audioInputStream);

        // Воспроизведение аудио
        clip.start();

        // Ждем, пока аудио не будет полностью воспроизведено
        while (clip.isRunning()) {
            Thread.sleep(10);
        }

        // Закрываем Clip и освобождаем ресурсы
        clip.close();
        audioInputStream.close();
        inputStream.close();

        // Отправка ответа об успешном прослушивании
        AudioServiceOuterClass.ListenResponse response = AudioServiceOuterClass.ListenResponse.newBuilder().setStatus("working....").build();

        responseStreamObserver.onNext(response);
        responseStreamObserver.onCompleted();
    } catch (IOException | UnsupportedAudioFileException | LineUnavailableException | InterruptedException e) {
        // Обработка исключений и отправка сообщения об ошибке
        String errorMessage = "Error occurred during audio playback: " + e.getMessage();
        AudioServiceOuterClass.ListenResponse errorResponse = AudioServiceOuterClass.ListenResponse.newBuilder().build();
        responseStreamObserver.onError(Status.INTERNAL.withDescription(errorMessage).asRuntimeException());
    }
}





    @Override
    public void upload(AudioServiceOuterClass.UploadRequest request, StreamObserver<AudioServiceOuterClass.UploadResponse> responseStreamObserver) {
        // Получение байтов аудио из запроса
        ByteString audioData = request.getAudioData();
        String nameId = request.getName();

        try {
            // Генерация пути и имени для выходного MP3-файла
            String fileId = nameId;  // Здесь должен быть ваш сгенерированный идентификатор файла
            String outputPath = "src/main/java/org/example/Book/" + fileId + ".mp3";

            // Запись байтов в файл MP3
            try (FileOutputStream outputStream = new FileOutputStream(outputPath)) {
                outputStream.write(audioData.toByteArray());
            }

            // Создание и заполнение объекта UploadResponse
            AudioServiceOuterClass.UploadResponse response = AudioServiceOuterClass.UploadResponse.newBuilder()
                    .setFileId("Вааау успешно скачали " + fileId)
                    .build();

            // Отправка ответа клиенту
            responseStreamObserver.onNext(response);
            responseStreamObserver.onCompleted();

            // Создание подключения к MongoDB
            MongoClientURI connectionString = new MongoClientURI("mongodb://atlas-sql-648b67ce6d0c983b4099ed12-lohh9.a.query.mongodb.net/name?ssl=true&authSource=admin");
            MongoClient mongoClient = new MongoClient(connectionString);

            // Получение базы данных
            MongoDatabase database = mongoClient.getDatabase("name");

            // Получение коллекции
            MongoCollection<Document> collection = database.getCollection("audio");

            // Создание документа для сохранения
            Document document = new Document();
            document.append("fileId", "1");
            document.append("fileName", fileId + ".mp3");
            document.append("fileLink", outputPath);

            // Вставка документа в коллекцию
//            collection.insertOne(document);



        } catch (IOException e) {
            // Обработка ошибки в случае возникновения проблем с записью файла
            responseStreamObserver.onError(e);
        }
    }



}
