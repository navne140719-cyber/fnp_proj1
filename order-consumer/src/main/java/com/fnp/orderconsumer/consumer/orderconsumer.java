package com.fnp.orderconsumer.consumer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fnp.orderconsumer.entity.order;
import com.fnp.orderconsumer.entity.Product;
import com.fnp.orderconsumer.repository.orderrepository;
import com.fnp.orderconsumer.repository.productrepository;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;
@Service
public class orderconsumer {
    private final productrepository productRepository;
    private final orderrepository orderRepository;
    private final ObjectMapper objectMapper = new ObjectMapper();
    public orderconsumer(productrepository productRepository, orderrepository orderRepository) {
        this.productRepository = productRepository;
        this.orderRepository = orderRepository;
    }
    @KafkaListener(topics = "order-created", groupId = "order-group")
    public void consume(String message) {
        try {
            System.out.println("Received Order Event: " + message);
            JsonNode json = objectMapper.readTree(message);
            Long orderId = json.get("orderId").asLong();
            JsonNode items = json.get("items");
            order orderObj = orderRepository.findById(orderId).orElse(null);
            if (orderObj == null) {
                System.out.println("Order not found");
                return;
            }
            boolean failed = false;
            for (JsonNode item : items) {
                Long productId = item.get("productId").asLong();
                int qty = item.get("qty").asInt();
                Product product = productRepository.findById(productId).orElse(null);
                if (product == null) {
                    System.out.println("Product not found: " + productId);
                    failed = true;
                    continue;
                }
                if (product.getStock() >= qty) {
                    product.setStock(product.getStock() - qty);
                    productRepository.save(product);
                }
                else {
                    System.out.println("Insufficient stock for product: " + productId);
                    failed = true;
                }
            }
            if(failed){
                orderObj.setStatus("FAILED");
            }
            else {
                orderObj.setStatus("COMPLETED");
            }
            orderRepository.save(orderObj);
            System.out.println("Order processing finished");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
