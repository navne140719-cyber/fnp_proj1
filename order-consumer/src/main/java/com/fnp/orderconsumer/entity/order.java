package com.fnp.orderconsumer.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;
import java.time.LocalDateTime;

@Entity
@Table(name = "orders")
@Getter
@Setter
public class order {

    @Id
    private Long id;

    private Long userId;

    private String status;

    private Double totalAmount;

    private LocalDateTime createdAt;
}
