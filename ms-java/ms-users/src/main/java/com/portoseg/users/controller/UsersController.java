package com.portoseg.users.controller;

import java.util.Optional;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.validation.Valid;

import org.apache.commons.lang3.StringUtils;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import com.portoseg.users.model.domain.UserDomain;
import com.portoseg.users.model.dto.request.CreateUserRequest;
import com.portoseg.users.model.dto.request.UpdateUserRequest;
import com.portoseg.users.repository.UsersRepository;
import com.portoseg.users.service.UsersService;

import lombok.RequiredArgsConstructor;

@RestController
@RequiredArgsConstructor
@RequestMapping("/v1/user")
public class UsersController {
    private final UsersService usersService;

    private final UsersRepository usersRepository;

    @PostMapping
    @ResponseStatus(HttpStatus.ACCEPTED)
    public void createUser(@RequestBody @Valid CreateUserRequest createUserRequest, HttpServletRequest request, HttpServletResponse response) {
        String userId = usersService.createUser(createUserRequest);
        String url = request.getRequestURL().toString();
        String location = StringUtils.join(url, "/", userId);

        response.addHeader("location", location);
    }

    @GetMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public Optional<UserDomain> getUser(@PathVariable String id) {
        return usersRepository.findById(id);
    }

    @PatchMapping("/{id}")
    @ResponseStatus(HttpStatus.OK)
    public UserDomain updateUser(@RequestBody @Valid UpdateUserRequest updateUserRequest, @PathVariable String id) {
        return usersService.updateUser(updateUserRequest, id);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteUser(@PathVariable String id) {
        usersService.deleteUser(id);
    }

}
