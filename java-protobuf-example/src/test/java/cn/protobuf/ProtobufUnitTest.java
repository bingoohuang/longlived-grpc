package cn.protobuf;


import org.junit.After;
import org.junit.Test;

import javax.xml.bind.DatatypeConverter;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Random;

import static org.junit.Assert.assertEquals;

public class ProtobufUnitTest {
    private final String filePath = "address_book";

    @After
    public void cleanup() throws IOException {
        Files.deleteIfExists(Paths.get(filePath));
    }

    @Test
    public void givenGeneratedProtobufClass_whenCreateClass_thenShouldCreateJavaInstance() {
        //when
        String email = "j@baeldung.com";
        int id = new Random().nextInt();
        String name = "Michael Program";
        String number = "01234567890";
        AddressBookProtos.Person p =
                AddressBookProtos.Person.newBuilder()
                        .setId(id)
                        .setName(name)
                        .setEmail(email)
                        .addNumbers(number)
                        .build();
        //then
        assertEquals(p.getEmail(), email);
        assertEquals(p.getId(), id);
        assertEquals(p.getName(), name);
        assertEquals(p.getNumbers(0), number);
    }


    @Test
    public void givenAddressBookWithOnePerson_whenSaveAsAFile_shouldLoadFromFileToJavaClass() throws IOException {
        //given
        String email = "j@baeldung.com";
//        int id = new Random().nextInt();
        int id = 1;
        String name = "Michael Program";
        String number = "01234567890";
        AddressBookProtos.Person p =
                AddressBookProtos.Person.newBuilder()
                        .setId(id)
                        .setName(name)
                        .setEmail(email)
                        .addNumbers(number)
                        .build();

        //when
        AddressBookProtos.AddressBook b0 = AddressBookProtos.AddressBook.newBuilder().addPeople(p).build();
        FileOutputStream fos = new FileOutputStream(filePath);
        b0.writeTo(fos);
        fos.close();

        byte[] data = b0.toByteArray();
        String b64 = DatatypeConverter.printBase64Binary(data);
        System.out.println(b64);

        data = DatatypeConverter.parseBase64Binary("CjAKD01pY2hhZWwgUHJvZ3JhbRABGg5qQGJhZWxkdW5nLmNvbSILMDEyMzQ1Njc4OTA=");
        AddressBookProtos.AddressBook b1 = AddressBookProtos.AddressBook.parseFrom(data);
        assertEquals(b0, b1);

        //then
        FileInputStream fis = new FileInputStream(filePath);
        AddressBookProtos.AddressBook b2 = AddressBookProtos.AddressBook.newBuilder().mergeFrom(fis).build();
        fis.close();
        p = b2.getPeople(0);
        assertEquals(p.getEmail(), email);
        assertEquals(p.getId(), id);
        assertEquals(p.getName(), name);
        assertEquals(p.getNumbers(0), number);
    }
}