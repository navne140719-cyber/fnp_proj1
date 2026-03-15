package com.fnp.orderconsumer.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Entity
@Table(name = "products")
@Getter
@Setter
public class Product {

    @Id
    private Long id;

    private String name;

    private Double price;

    private Integer stock;
}