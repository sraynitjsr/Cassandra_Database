package com.sraynitjsr.cassandrajavaintegration;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.data.cassandra.core.mapping.Column;
import org.springframework.data.cassandra.core.mapping.Table;
import org.springframework.data.cassandra.repository.CassandraRepository;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@SpringBootApplication
public class JavaBasedCassandraAPI {

    public static void main(String[] args) {
        SpringApplication.run(JavaBasedCassandraAPI.class, args);
    }

    @Table("users")
    public static class User {

        @Column("id")
        private UUID id;

        @Column("first_name")
        private String firstName;

        @Column("last_name")
        private String lastName;

        @Column("email")
        private String email;

        public User() {}

        public User(UUID id, String firstName, String lastName, String email) {
            this.id = id;
            this.firstName = firstName;
            this.lastName = lastName;
            this.email = email;
        }

        public UUID getId() {
            return id;
        }

        public void setId(UUID id) {
            this.id = id;
        }

        public String getFirstName() {
            return firstName;
        }

        public void setFirstName(String firstName) {
            this.firstName = firstName;
        }

        public String getLastName() {
            return lastName;
        }

        public void setLastName(String lastName) {
            this.lastName = lastName;
        }

        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }
    }

    public interface UserRepository extends CassandraRepository<User, UUID> {
    }

    @RestController
    @RequestMapping("/api/users")
    public static class UserController {

        @Autowired
        private UserRepository userRepository;

        @GetMapping
        public List<User> getAllUsers() {
            return userRepository.findAll();
        }

        @GetMapping("/{id}")
        public Optional<User> getUserById(@PathVariable UUID id) {
            return userRepository.findById(id);
        }

        @PostMapping
        public User createUser(@RequestBody User user) {
            user.setId(UUID.randomUUID());
            return userRepository.save(user);
        }

        @PutMapping("/{id}")
        public User updateUser(@PathVariable UUID id, @RequestBody User user) {
            if (userRepository.existsById(id)) {
                user.setId(id);
                return userRepository.save(user);
            }
            return null;
        }

        @DeleteMapping("/{id}")
        public void deleteUser(@PathVariable UUID id) {
            userRepository.deleteById(id);
        }
    }

    @Autowired
    private UserRepository userRepository;

    @Bean
    public CommandLineRunner demo() {
        return (args) -> {
            userRepository.save(new User(UUID.randomUUID(), "Aarav", "Sharma", "aarav.sharma@example.com"));
            userRepository.save(new User(UUID.randomUUID(), "Ishita", "Patel", "ishita.patel@example.com"));
            userRepository.save(new User(UUID.randomUUID(), "Saanvi", "Reddy", "saanvi.reddy@example.com"));
            userRepository.save(new User(UUID.randomUUID(), "Aditya", "Verma", "aditya.verma@example.com"));
            userRepository.save(new User(UUID.randomUUID(), "Ananya", "Gupta", "ananya.gupta@example.com"));
        };
    }
}
