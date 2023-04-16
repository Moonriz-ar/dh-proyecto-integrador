import { useRouter } from "next/router";
import Link from "next/link";
import { useForm } from "react-hook-form";

import Button from "@/components/button";

type Inputs = {
  email: string;
  password: string;
};

function Login() {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Inputs>();

  const onSubmit = (data: Inputs) => {
    console.log(data);
    router.push("/");
  };

  return (
    <section className="flex flex-col justify-center md:h-full md:max-w-sm md:mx-auto">
      <h1 className="mb-5 text-2xl font-bold text-center text-primary">
        Iniciar sesión
      </h1>
      <form
        onSubmit={handleSubmit(onSubmit)}
        // onSubmit={e => {e.preventDefault}}
        className="flex flex-col gap-5 mb-5"
      >
        <div className="flex flex-col gap-2">
          <label htmlFor="email" className="text-sm font-medium text-secondary">
            Correo Electrónico
          </label>
          <input
            id="email"
            className="h-10 p-3 rounded shadow"
            type="text"
            {...register("email", {
              required: { value: true, message: "Campo obligatorio" },
              pattern: {
                value:
                  /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                message: "Ingrese un mail válido",
              },
            })}
          />
          {errors.email && (
            <p className="text-xs font-medium text-red-700">
              {errors.email.message}
            </p>
          )}
        </div>
        <div className="flex flex-col gap-2">
          <label
            htmlFor="password"
            className="text-sm font-medium text-secondary"
          >
            Contraseña
          </label>
          <input
            id="password"
            type="password"
            className="h-10 p-3 rounded shadow"
            {...register("password", {
              required: { value: true, message: "Campo obligatorio" },
              minLength: {
                value: 6,
                message: "La contraseña debe tener al menos 6 caracteres",
              },
            })}
          />
          {errors.password && (
            <p className="text-xs font-medium text-red-700">
              {errors.password.message}
            </p>
          )}
        </div>

        <Button variant="contained">Ingresar</Button>
      </form>
      <p className="text-xs text-center">
        ¿Aún no tenes cuenta?{" "}
        <Link href="/signup" className="text-cyan-600">
          Registrate
        </Link>
      </p>
    </section>
  );
}

export default Login;
