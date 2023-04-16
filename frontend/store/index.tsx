import { proxy, useSnapshot, subscribe } from "valtio";
import { devtools } from "valtio/utils";

import { User } from "@/types";

type NullableUser = User | null;

let storedUser;

// fix when Next.js try to access localStorage before the browser renders it
if (typeof window !== "undefined") {
  // 👉️ can use localStorage here
  storedUser = localStorage.getItem("user");
} else {
  console.log("Can't use localStorage");
  // 👉️ can't use localStorage
}

const store = proxy<{ user: NullableUser }>(
  storedUser ? JSON.parse(storedUser) : { user: null }
);

const unsub = devtools(store, { name: "store", enabled: true });

subscribe(store, () => {
  localStorage.setItem("user", JSON.stringify(store));
});

const setUser = (user: NullableUser) => {
  store.user = user;
};

export { store, useSnapshot, subscribe, setUser };
