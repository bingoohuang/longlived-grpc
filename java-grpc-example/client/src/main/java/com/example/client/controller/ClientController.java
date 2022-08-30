package com.example.client.controller;

import cn.hutool.core.lang.Pair;
import cn.hutool.core.util.RandomUtil;
import com.example.client.client.SimpleClient;
import java.util.List;
import java.util.stream.Collectors;
import lombok.var;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import simple.testgrpc.Simple.SimpleResponse;

@RestController
@RequestMapping("/client")
public class ClientController {

  @Autowired
  private SimpleClient simpleClient;

  @GetMapping
  public Pair<Long, String> rPCRequest() {
    var response = simpleClient.rPCRequest(RandomUtil.randomLong(), RandomUtil.randomString(16));
    return toPair(response);
  }

  @GetMapping("/clientStreaming")
  public Pair<Long, String> clientStreaming(@RequestParam(required = false) Integer count) {
    if (count == null || count.intValue() <= 0) {
      count = 3;
    }
    var response = simpleClient.clientStreaming(count);
    return toPair(response);
  }

  @GetMapping("/serverStreaming")
  public List<Pair<Long, String>> serverStreaming() {
    var response = simpleClient.serverStreaming(RandomUtil.randomLong(),
        RandomUtil.randomString(16));
    return response.stream().map(this::toPair).collect(Collectors.toList());
  }

  @GetMapping("/streamingBiDirectional")
  public List<Pair<Long, String>> streamingBiDirectional(
      @RequestParam(required = false) Integer count) {
    if (count == null || count.intValue() <= 0) {
      count = 3;
    }
    var response = simpleClient.streamingBiDirectional(count);
    return response.stream().map(this::toPair).collect(Collectors.toList());
  }

  private Pair<Long, String> toPair(SimpleResponse response) {
    return Pair.of(response.getResponseId(), response.getResponse());
  }

}
