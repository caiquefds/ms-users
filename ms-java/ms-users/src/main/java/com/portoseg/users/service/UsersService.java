package com.portoseg.users.service;

import java.util.Optional;

import org.apache.commons.lang3.StringUtils;
import org.springframework.stereotype.Service;

import com.portoseg.users.exception.UnprocessableEntityException;
import com.portoseg.users.model.domain.UserDomain;
import com.portoseg.users.model.dto.request.CreateUserRequest;
import com.portoseg.users.model.dto.request.UpdateUserRequest;
import com.portoseg.users.repository.UsersRepository;

import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Service
@RequiredArgsConstructor
public class UsersService {

    private final UsersRepository usersRepository;

    public String createUser(CreateUserRequest createUserRequest) {
        log.info("Creating User");

        UserDomain userDomain = UserDomain.valueOf(createUserRequest);

        boolean isPresent = usersRepository.findByEmailAndUsername(userDomain.getEmail(), userDomain.getUsername()).isPresent();

        if (isPresent) {
            log.error("This user already exists");
            throw new UnprocessableEntityException("422.000");
        }

        UserDomain userDomainCreated = usersRepository.save(userDomain);
        return userDomainCreated.getId();
    }

    public UserDomain updateUser(UpdateUserRequest updateUserRequest, String id) {
        UserDomain userDomain = usersRepository.findById(id)
                .orElseThrow(() -> {
                    log.info("User id: {} does not exist", id);
                    throw new UnprocessableEntityException("422.001");
                });

        Optional.ofNullable(updateUserRequest.getEmail())
                .filter(StringUtils::isNotBlank)
                .ifPresent(userDomain::setEmail);

        Optional.ofNullable(updateUserRequest.getUsername())
                .ifPresent(userDomain::setUsername);

        Optional.ofNullable(updateUserRequest.getPassword())
                .ifPresent(userDomain::setPassword);

        Optional.ofNullable(updateUserRequest.getStatus())
                .ifPresent(userDomain::setStatus);

        usersRepository.save(userDomain);
        log.info("User with id {} updated", id);
        return userDomain;
    }

    public void deleteUser(String id) {
        log.info("Deleting User by id: {}...", id);
        usersRepository.findById(id).ifPresent(usersRepository::delete);
    }

}
