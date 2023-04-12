package com.portoseg.users.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.portoseg.users.model.domain.UserDomain;

@Repository
public interface UsersRepository extends JpaRepository<UserDomain, String> {

    Optional<UserDomain> findByEmailAndUsername(String email, String username);

}
