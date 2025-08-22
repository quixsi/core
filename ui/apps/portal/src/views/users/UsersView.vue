<template>
    <div>
        <h1 class="text-2xl font-bold mb-4">Users</h1>
        <Card class="max-w-2xl mx-auto mt-8">
        <template #content>
            <div v-if="loading" class="mb-4">Loading users...</div>
            <div v-else-if="error" class="mb-4 text-red-500">{{ error }}</div>
            <table v-if="users.length" class="min-w-full">
            <thead>
                <tr>
                <th class="px-4 py-2 border-b">Name</th>
                <th class="px-4 py-2 border-b">Email</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in users" :key="user.id">
                    <td class="px-4 py-2 border-b flex items-center gap-2">
                        {{ user.name }}
                        <button 
                            @click="deleteUser(user.id)" 
                            title="Delete user" 
                            class="text-red-500 hover:text-red-700 cursor-pointer"
                        >
                            <svg 
                                xmlns="http://www.w3.org/2000/svg" 
                                class="inline w-4 h-4" 
                                fill="none" 
                                viewBox="0 0 24 24" 
                                stroke="currentColor"
                            >
                                <path 
                                    stroke-linecap="round" 
                                    stroke-linejoin="round" 
                                    stroke-width="2" 
                                    d="M6 18L18 6M6 6l12 12"
                                />
                            </svg>
                        </button>
                    </td>
                    <td class="px-4 py-2 border-b">{{ user.email }}</td>
                </tr>
            </tbody>
            </table>
            <div v-else class="text-gray-500">No users found.</div>
        </template>
        </Card>
        <CreateUser />
    </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { Card } from '@quixsi/components'
import { CreateUser } from './create'
import { useUsers } from './useUsers'
const { users, loading, error, fetchUsers, deleteUser } = useUsers()

onMounted(fetchUsers)
</script>