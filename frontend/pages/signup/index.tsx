import { useRouter } from "next/router";
import Link from "next/link";
import { useForm } from "react-hook-form";

import Button from "@/components/button";

type Inputs = {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  confirmPassword: string;
};

function Signup() {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors },
  } = useForm<Inputs>();

  const onSubmit = (data: Inputs) => {
    console.log(data);
    router.push("/");
  };

  return (
    <section className="flex flex-col justify-center md:h-full md:max-w-sm md:mx-auto">
      <h1 className="mb-5 text-2xl font-bold text-center text-primary">
        Crear cuenta
      </h1>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col gap-5 mb-5"
      >
        <div className="flex flex-col gap-2">
          <label
            htmlFor="firstName"
            className="text-sm font-medium text-secondary"
          >
            Nombre
          </label>
          <input
            id="firstName"
            className="h-10 p-3 rounded shadow"
            type="text"
            {...register("firstName", {
              required: { value: true, message: "Campo obligatorio" },
            })}
          />
          {errors.firstName && (
            <p className="text-xs font-medium text-red-700">
              {errors.firstName.message}
            </p>
          )}
        </div>
        <div className="flex flex-col gap-2">
          <label
            htmlFor="lastName"
            className="text-sm font-medium text-secondary"
          >
            Apellido
          </label>
          <input
            id="lastName"
            className="h-10 p-3 rounded shadow"
            type="text"
            {...register("lastName", {
              required: { value: true, message: "Campo obligatorio" },
            })}
          />
          {errors.lastName && (
            <p className="text-xs font-medium text-red-700">
              {errors.lastName.message}
            </p>
          )}
        </div>
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
        <div className="flex flex-col gap-2">
          <label
            htmlFor="confirmPassword"
            className="text-sm font-medium text-secondary"
          >
            Confirmar contraseña
          </label>
          <input
            id="confirmPassword"
            type="password"
            className="h-10 p-3 rounded shadow"
            {...register("confirmPassword", {
              required: { value: true, message: "Campo obligatorio" },
              validate: (val: string) => {
                if (watch("password") != val) {
                  return "La contraseña no coincide";
                }
              },
              minLength: {
                value: 6,
                message: "La contraseña debe tener al menos 6 caracteres",
              },
            })}
          />
          {errors.confirmPassword && (
            <p className="text-xs font-medium text-red-700">
              {errors.confirmPassword.message}
            </p>
          )}
        </div>

        <Button variant="contained">Ingresar</Button>
      </form>
      <p className="text-xs text-center">
        ¿Ya tienes una cuenta?{" "}
        <Link href="/login" className="text-cyan-600">
          Iniciar sesión
        </Link>
      </p>
    </section>
  );
}

export default Signup;
