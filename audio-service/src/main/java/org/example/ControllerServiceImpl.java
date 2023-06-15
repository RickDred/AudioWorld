package org.example;
import audio.AudioServiceGrpc;
import audio.AudioServiceOuterClass;
import com.google.protobuf.ByteString;
import com.mongodb.MongoClient;
import com.mongodb.MongoClientURI;
import com.mongodb.client.MongoCollection;
import com.mongodb.client.MongoDatabase;
import io.grpc.stub.StreamObserver;
import org.bson.Document;
import java.io.FileOutputStream;
import java.io.IOException;

public class ControllerServiceImpl extends AudioServiceGrpc.AudioServiceImplBase {
    // Создание подключения к MongoDB
//    MongoClientURI connectionString = new MongoClientURI("mongodb://atlas-sql-648b67ce6d0c983b4099ed12-lohh9.a.query.mongodb.net/name?ssl=true&authSource=admin");
//    MongoClient mongoClient = new MongoClient(connectionString);
//
//    // Получение базы данных
//    MongoDatabase database = mongoClient.getDatabase("name");


//
    @Override
    public void audioListen(AudioServiceOuterClass.ListenRequest request, StreamObserver<AudioServiceOuterClass.ListenResponse> responseStreamObserver)
    {
        System.out.println(request);

        AudioServiceOuterClass.ListenResponse response = AudioServiceOuterClass.ListenResponse.newBuilder().setStatus(0).build();

        responseStreamObserver.onNext(response);
        responseStreamObserver.onCompleted();
    }

    @Override
    public void Upload(AudioServiceOuterClass.UploadRequest request, StreamObserver<AudioServiceOuterClass.UploadResponse> responseStreamObserver) {

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
