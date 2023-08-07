package com.example.pingpong.controller;

import com.example.pingpong.service.PingPongService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HomeController {

    private PingPongService pingPongService;

    private HomeController(PingPongService pingPongService) {
        this.pingPongService = pingPongService;
    }

    @GetMapping("/")
    public String home() {
        return "Total ping-pong count: " + pingPongService.getCount() + "\n";
    }

    @GetMapping("/hello")
    public String hello() {
        return "Hello Docker, from Spring App!";
    }
}
