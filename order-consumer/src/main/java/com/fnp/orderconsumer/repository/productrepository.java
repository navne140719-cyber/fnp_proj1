package com.fnp.orderconsumer.repository;

import com.fnp.orderconsumer.entity.Product;
import org.springframework.data.jpa.repository.JpaRepository;

public interface productrepository extends JpaRepository<Product, Long> {

}