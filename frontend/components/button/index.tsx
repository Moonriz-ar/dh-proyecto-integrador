import styles from "./button.module.css";
import { clsx } from "clsx";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "outlined" | "contained" | "text";
  children?: React.ReactNode;
}

const Button: React.FC<ButtonProps> = ({ children, variant, ...props }) => {
  return (
    <button
      className={clsx({
        [styles.button]: true,
        [styles.contained]: variant === "contained" || variant === undefined,
        [styles.outlined]: variant === "outlined",
        [styles.text]: variant === "text",
      })}
      {...props}
    >
      {children}
    </button>
  );
};

export default Button;
