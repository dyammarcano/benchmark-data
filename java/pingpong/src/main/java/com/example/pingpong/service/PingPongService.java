package com.example.pingpong.service;

import org.springframework.stereotype.Service;

@Service
public class PingPongService {
    private byte[] buffer;
    private final int maxBufferSize = 1000 * 1024 * 1024; // 100MB
    private int count = 0;

    public PingPongService() {
        new Thread(() -> {
            while (count <= 20000) {
                try {
                    // Generate load by repeatedly creating buffers up to 100MB and resetting to 0MB
                    if (buffer == null || buffer.length >= maxBufferSize) {
                        buffer = new byte[1 * 1024 * 1024]; // 1MB
                    }
                    buffer = new byte[buffer.length + 1 * 1024 * 1024]; // 1MB
                    count++;
                    Thread.sleep(10);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
        }).start();
    }

    public int getCount() {
        return count;
    }
}
