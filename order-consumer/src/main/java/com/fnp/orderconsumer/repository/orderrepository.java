package com.fnp.orderconsumer.repository;

import com.fnp.orderconsumer.entity.order;
import org.springframework.data.jpa.repository.JpaRepository;

public interface orderrepository extends JpaRepository<order, Long> {

}